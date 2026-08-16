// Package filecache provides a high-performance, thread-safe, in-memory cache for local disk files
// and remote HTTPS resources (e.g., Starlark rule files, OpenID configurations, and JSON schema policies).
//
// Key Design Principles for Maintainers:
//
// 1. Dual Resource Support:
//   - Local Disk Files: Revalidated using os.Stat modification times. If a disk file hasn't changed,
//     the cache extends its TTL without re-reading the file from disk.
//   - Remote HTTPS Resources: Revalidated using HTTP ETag (If-None-Match) and Expires headers. Supports
//     HTTP 304 Not Modified to refresh TTLs without re-downloading response bodies.
//
// 2. Concurrency & Data Race Safety (Immutability):
//   - FileEntry pointers stored in sync.Map are treated as strictly IMMUTABLE.
//   - Whenever a cache entry is refreshed or revalidated, a new FileEntry struct copy is allocated
//     and stored into sync.Map. Never mutate fields of an existing FileEntry pointer in-place.
//
// 3. Thundering Herd Suppression (Single-Flight Coalescing):
//   - Uses golang.org/x/sync/singleflight to coalesce concurrent cache misses or expiration fetches.
//   - If 100 concurrent requests request an expired resource simultaneously, only 1 network fetch or
//     disk read is executed, and all 100 callers safely share the result.
//
// 4. Fallback on Error:
//   - If a remote server is unreachable during an expired entry revalidation, the cache returns the
//     stale cached entry so the system can keep operating gracefully and eventually update the file.
package filecache

import "time"

// FileCache abstracts caching operations for disk files and HTTPS URLs.
type FileCache interface {
	// Get retrieves a file or HTTPS resource from the cache, revalidating it if expired.
	Get(fileName string) (*FileEntry, error)

	// Set explicitly stores file contents in the cache with the specified TTL duration.
	Set(fileName string, content []byte, ttl time.Duration) error

	// MustExist returns an entry directly from the cache without checking expiration or revalidating.
	MustExist(fileName string) (*FileEntry, error)
}
