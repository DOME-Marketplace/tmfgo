// Package filecache implements a simple file cache for local and remote files.
// It supports caching of file contents in memory, with configurable freshness
// policies for both disk-based and server-based files. It handles HTTP ETag
// and Expires headers to optimize network usage and ensure data consistency.
package filecache

import (
	"errors"
	"fmt"
	"hash/maphash"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/DOME-Marketplace/tmfgo/internal/errl"
	"gitlab.com/greyxor/slogor"
	"golang.org/x/sync/singleflight"
)

// FileEntry represents a cached resource (local disk file or remote HTTPS URL).
//
// MAINTAINER NOTE ON THREAD SAFETY & IMMUTABILITY:
// FileEntry instances stored in sync.Map MUST BE TREATED AS IMMUTABLE once published.
// Never mutate fields of an existing FileEntry pointer in-place! When updating an entry's
// TTL or headers, construct and store a new FileEntry struct copy. This guarantees that
// concurrent readers accessing the cached pointer will never experience data races.
type FileEntry struct {
	Name         string    // Cache key (absolute file path or HTTPS URL)
	EntryUpdated time.Time // Timestamp when this entry was created or revalidated in cache
	FileModTime  time.Time // Last modification time of the file (from os.Stat or HTTP Last-Modified)
	Content      []byte    // Cached raw file bytes
	Etag         string    // HTTP ETag header value used for conditional GET (If-None-Match)
	Expires      time.Time // Expiration deadline. Until now.Before(Expires), cache hits return immediately
	FileHash     uint64    // maphash hash of Content for quick byte-level change detection
}

const maxFileSize = 1024 * 1024                 // Default maximum allowed file size (1MB)
const freshnessForDiskFiles = 20 * time.Second  // Default TTL for disk files before checking os.Stat
const freshnessForServerFiles = 5 * time.Minute // Default TTL for HTTPS URLs when no Expires header is returned

// ErrorRedirectsNotAllowed is returned by the HTTP client if a remote server attempts a redirect.
var ErrorRedirectsNotAllowed = errors.New("redirects not allowed")

// Seed used for maphash byte hashing.
var seed = maphash.MakeSeed()

// SimpleFileCache implements the FileCache interface.
// It uses sync.Map for concurrent lock-free reads and singleflight.Group to suppress
// thundering herd cache stampedes when revalidating expired resources.
type SimpleFileCache struct {
	fileCache  *sync.Map          // Thread-safe map: string (fileName/URL) -> *FileEntry
	httpClient *http.Client       // HTTP client used for remote URL fetches
	options    *FileCacheOptions  // Configuration settings (TTLs, max size, HTTP client)
	sfGroup    singleflight.Group // Single-flight coalescer to prevent thundering herd fetches
}

// FileCacheOptions configures the behavior of SimpleFileCache.
type FileCacheOptions struct {
	// MaxFileSize is the maximum allowed size of a cached file in bytes (Default: 1MB).
	MaxFileSize int
	// FreshnessForDiskFiles is the TTL for disk files before checking os.Stat (Default: 20s).
	FreshnessForDiskFiles time.Duration
	// FreshnessForServerFiles is the default TTL for HTTPS URLs when no Expires header is present (Default: 5min).
	FreshnessForServerFiles time.Duration
	// HTTPClient is the custom HTTP client used for remote requests (Default: 10s timeout, no redirects).
	HTTPClient *http.Client
}

// NewSimpleFileCache creates and initializes a new SimpleFileCache with the given options.
// If options is nil, sensible defaults are applied:
// - MaxFileSize: 1MB
// - FreshnessForDiskFiles: 20 seconds
// - FreshnessForServerFiles: 5 minutes
// - HTTPClient: 10-second timeout, redirects forbidden
func NewSimpleFileCache(options *FileCacheOptions) *SimpleFileCache {
	if options == nil {
		options = &FileCacheOptions{}
	}
	if options.MaxFileSize == 0 {
		options.MaxFileSize = maxFileSize
	}
	if options.FreshnessForDiskFiles == 0 {
		options.FreshnessForDiskFiles = freshnessForDiskFiles
	}
	if options.FreshnessForServerFiles == 0 {
		options.FreshnessForServerFiles = freshnessForServerFiles
	}

	var client *http.Client
	if options.HTTPClient == nil {
		client = &http.Client{
			Timeout: 10 * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return ErrorRedirectsNotAllowed
			},
		}
	} else {
		client = options.HTTPClient
	}

	return &SimpleFileCache{
		fileCache:  &sync.Map{},
		httpClient: client,
		options:    options,
	}
}

// Get is the main entry point for retrieving cached resources.
// It inspects the resource string prefix:
// - Strings starting with "https://" are dispatched to GetURL.
// - All other strings are treated as local disk file paths and dispatched to GetFile.
func (m *SimpleFileCache) Get(fileName string) (*FileEntry, error) {
	if fileName == "" {
		return nil, errl.Errorf("file name is empty")
	}

	if strings.HasPrefix(fileName, "https://") {
		return m.GetURL(fileName)
	}
	return m.GetFile(fileName)
}

// GetURL retrieves a remote HTTPS resource, utilizing HTTP ETags, Expires headers, and singleflight coalescing.
//
// Execution Flow for Maintainers:
// 1. FAST PATH: Perform an atomic sync.Map lookup. If found and now.Before(Expires), return immediately.
// 2. SLOW PATH: If missing or expired, use singleflight.Group to ensure only ONE goroutine executes the network fetch.
// 3. Inside Singleflight:
//   - Re-check cache in case a parallel singleflight call completed while waiting.
//   - Call fetchURL with the existing entry to send a conditional HTTP GET (If-None-Match).
//   - On HTTP 304 Not Modified: revalidate existing entry TTL/headers without re-downloading content.
//   - On HTTP 200 OK: parse new content, calculate hash, store new FileEntry in cache.
//   - On Network Error: if a stale cached entry exists, log the error and return the stale entry
//     so the application remains available even during network outages.
func (m *SimpleFileCache) GetURL(fileName string) (*FileEntry, error) {
	now := time.Now()

	// Step 1: Fast Path — Atomic cache lookup
	fe, found := m.fileCache.Load(fileName)
	var existing *FileEntry
	if found {
		existing = fe.(*FileEntry)
		if now.Before(existing.Expires) {
			slog.Debug("file cache entry is fresh", slog.String("file", fileName))
			return existing, nil
		}
		slog.Debug("file cache entry expired, revalidating", slog.String("file", fileName))
	} else {
		slog.Debug("file cache entry not found, fetching from server", slog.String("file", fileName))
	}

	// Step 2: Slow Path — Singleflight request coalescing to suppress thundering herd
	val, err, shared := m.sfGroup.Do(fileName, func() (any, error) {
		// Double-check cache in case a prior singleflight fetch finished while this goroutine waited
		if fe, found := m.fileCache.Load(fileName); found {
			entry := fe.(*FileEntry)
			if time.Now().Before(entry.Expires) {
				return entry, nil
			}
		}

		// Perform network fetch / conditional HTTP GET
		entry, err := m.fetchURL(fileName, existing)
		if err != nil {
			// Fallback: If network fetch fails but we have stale cached data, serve stale data
			if existing != nil {
				slog.Error("error revalidating URL file, returning stale cached entry", slog.String("file", fileName), slogor.Err(err))
				return existing, nil
			}
			return nil, err
		}

		m.fileCache.Store(fileName, entry)
		return entry, nil
	})

	if err != nil {
		return nil, err
	}

	if shared {
		slog.Debug("thundering herd suppressed for URL fetch", slog.String("file", fileName))
	}

	return val.(*FileEntry), nil
}

// fetchURL performs the actual HTTP GET request to the remote URL.
func (m *SimpleFileCache) fetchURL(url string, existing *FileEntry) (*FileEntry, error) {
	now := time.Now()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errl.Errorf("error creating HTTP request for %s: %w", url, err)
	}

	// Add conditional request header if we have a cached ETag
	if existing != nil && existing.Etag != "" {
		req.Header.Add("If-None-Match", existing.Etag)
	}

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return nil, errl.Errorf("error fetching %s: %w", url, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// Handle HTTP 304 Not Modified (Server indicates content hasn't changed)
	if resp.StatusCode == http.StatusNotModified {
		if existing == nil {
			return nil, errl.Errorf("received 304 Not Modified for %s but no existing entry in cache", url)
		}
		slog.Debug("HTTP 304 Not Modified", slog.String("file", url))
		return m.revalidateURLEntry(existing, resp, now), nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("readFileIfNew: error reading file %s: status %d", url, resp.StatusCode)
	}

	// Read response body with io.LimitReader to prevent OOM on large files
	limitReader := io.LimitReader(resp.Body, int64(m.options.MaxFileSize+1))
	content, err := io.ReadAll(limitReader)
	if err != nil {
		return nil, errl.Errorf("error reading response body for %s: %w", url, err)
	}

	if len(content) > m.options.MaxFileSize {
		return nil, fmt.Errorf("readFileIfNew: file %s is too big", url)
	}

	return m.newURLEntry(url, content, resp, now), nil
}

// revalidateURLEntry creates a fresh FileEntry copy with updated ETag/Expires headers on HTTP 304.
func (m *SimpleFileCache) revalidateURLEntry(existing *FileEntry, resp *http.Response, now time.Time) *FileEntry {
	if existing == nil {
		return nil
	}

	// IMMUTABILITY: Allocate a new struct copy to prevent mutating existing pointer in-place
	updated := *existing
	updated.EntryUpdated = now

	if etag := resp.Header.Get("Etag"); etag != "" {
		updated.Etag = etag
	}

	if expires := resp.Header.Get("Expires"); expires != "" {
		if t, err := http.ParseTime(expires); err == nil {
			updated.Expires = t
		}
	}

	if lastMod := resp.Header.Get("Last-Modified"); lastMod != "" {
		if t, err := http.ParseTime(lastMod); err == nil {
			updated.FileModTime = t
		}
	}

	return &updated
}

// newURLEntry constructs a brand-new FileEntry from an HTTP 200 response.
func (m *SimpleFileCache) newURLEntry(url string, content []byte, resp *http.Response, now time.Time) *FileEntry {
	entry := &FileEntry{
		Name:         url,
		EntryUpdated: now,
		Content:      content,
		FileHash:     maphash.Bytes(seed, content),
	}

	if etag := resp.Header.Get("Etag"); etag != "" {
		entry.Etag = etag
	}

	if lastMod := resp.Header.Get("Last-Modified"); lastMod != "" {
		if t, err := http.ParseTime(lastMod); err == nil {
			entry.FileModTime = t
		}
	}

	// Parse Expires header if provided, otherwise apply default server file TTL
	if expires := resp.Header.Get("Expires"); expires != "" {
		if t, err := http.ParseTime(expires); err == nil {
			entry.Expires = t
		} else {
			entry.Expires = now.Add(m.options.FreshnessForServerFiles)
		}
	} else {
		entry.Expires = now.Add(m.options.FreshnessForServerFiles)
	}

	return entry
}

// GetFile retrieves a local disk file from cache, revalidating its modification time via os.Stat when expired.
//
// Execution Flow for Maintainers:
// 1. FAST PATH: Perform atomic sync.Map lookup. If found and now.Before(Expires), return immediately.
// 2. SLOW PATH: If missing or expired, use singleflight.Group to coalesce concurrent disk stat/read requests.
// 3. Inside Singleflight:
//   - Re-check cache in case a parallel fetch finished while waiting.
//   - Execute os.Stat(fileName) to inspect disk modification time.
//   - If file on disk has NOT changed since existing.FileModTime: extend TTL without re-reading file content.
//   - If file on disk WAS modified or is missing from cache: call os.ReadFile(fileName) and store new FileEntry.
func (m *SimpleFileCache) GetFile(fileName string) (*FileEntry, error) {
	now := time.Now()

	// Step 1: Fast Path — Atomic cache lookup
	fe, found := m.fileCache.Load(fileName)
	var existing *FileEntry
	if found {
		existing = fe.(*FileEntry)
		if now.Before(existing.Expires) {
			slog.Debug("file cache entry is fresh", slog.String("file", fileName))
			return existing, nil
		}
		slog.Debug("file cache entry expired, checking disk modification time", slog.String("file", fileName))
	} else {
		slog.Debug("file cache entry not found, reading from disk", slog.String("file", fileName))
	}

	// Step 2: Slow Path — Singleflight coalescing for disk file stat/read
	val, err, shared := m.sfGroup.Do(fileName, func() (any, error) {
		if fe, found := m.fileCache.Load(fileName); found {
			entry := fe.(*FileEntry)
			if time.Now().Before(entry.Expires) {
				return entry, nil
			}
		}

		entry, err := m.fetchFile(fileName, existing)
		if err != nil {
			return nil, err
		}

		m.fileCache.Store(fileName, entry)
		return entry, nil
	})

	if err != nil {
		return nil, err
	}

	if shared {
		slog.Debug("thundering herd suppressed for disk file read", slog.String("file", fileName))
	}

	return val.(*FileEntry), nil
}

// fetchFile performs os.Stat and os.ReadFile operations for disk files.
func (m *SimpleFileCache) fetchFile(fileName string, existing *FileEntry) (*FileEntry, error) {
	now := time.Now()

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return nil, errl.Errorf("readFileIfNew: error checking file %s: %w", fileName, err)
	}
	if fileInfo.Mode().IsDir() {
		return nil, errl.Errorf("readFileIfNew: file %s is a directory, not a file", fileName)
	}

	if fileInfo.Size() > int64(m.options.MaxFileSize) {
		return nil, errl.Errorf("readFileIfNew: file %s is too big", fileName)
	}

	modifiedAt := fileInfo.ModTime()

	// OPTIMIZATION: If entry exists and file modtime on disk hasn't changed, extend TTL without re-reading file content
	if existing != nil && !existing.FileModTime.Before(modifiedAt) {
		slog.Debug("file not modified on disk, extending TTL", slog.String("file", fileName))
		// IMMUTABILITY: Create a fresh entry copy
		updated := *existing
		updated.EntryUpdated = now
		updated.Expires = now.Add(m.options.FreshnessForDiskFiles)
		return &updated, nil
	}

	// File was modified or not previously cached, read content from disk
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errl.Errorf("readFileIfNew: error reading file %s: %w", fileName, err)
	}

	return &FileEntry{
		Name:         fileName,
		EntryUpdated: now,
		FileModTime:  modifiedAt,
		Content:      content,
		FileHash:     maphash.Bytes(seed, content),
		Expires:      now.Add(m.options.FreshnessForDiskFiles),
	}, nil
}

// Set explicitly stores file contents in the cache with the given TTL duration.
// If ttl is 0, the entry is set with a 100-year default TTL (effectively non-expiring).
func (m *SimpleFileCache) Set(fileName string, content []byte, ttl time.Duration) error {
	if fileName == "" {
		return errl.Errorf("file name is empty")
	}

	now := time.Now()
	if ttl == 0 {
		ttl = time.Hour * 24 * 365 * 100 // 100 years
	}

	entry := &FileEntry{
		Name:         fileName,
		EntryUpdated: now,
		FileModTime:  now,
		Content:      content,
		FileHash:     maphash.Bytes(seed, content),
		Expires:      now.Add(ttl),
	}

	m.fileCache.Store(fileName, entry)
	return nil
}

// MustExist returns an entry directly from the cache without checking expiration or revalidating.
// If the key is missing from cache, it returns a non-nil error.
func (m *SimpleFileCache) MustExist(fileName string) (*FileEntry, error) {
	fe, found := m.fileCache.Load(fileName)
	if !found {
		return nil, errl.Errorf("file %s not found in cache", fileName)
	}
	return fe.(*FileEntry), nil
}
