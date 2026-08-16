package filecache

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSimpleFileCache(t *testing.T) {
	// Test default options
	cache := NewSimpleFileCache(nil)
	assert.NotNil(t, cache)
	assert.Equal(t, maxFileSize, cache.options.MaxFileSize)
	assert.Equal(t, freshnessForDiskFiles, cache.options.FreshnessForDiskFiles)
	assert.Equal(t, freshnessForServerFiles, cache.options.FreshnessForServerFiles)
	assert.NotNil(t, cache.httpClient)

	// Test custom options
	customOpts := &FileCacheOptions{
		MaxFileSize:             500,
		FreshnessForDiskFiles:   time.Minute,
		FreshnessForServerFiles: time.Hour,
	}
	cache = NewSimpleFileCache(customOpts)
	assert.Equal(t, 500, cache.options.MaxFileSize)
	assert.Equal(t, time.Minute, cache.options.FreshnessForDiskFiles)
	assert.Equal(t, time.Hour, cache.options.FreshnessForServerFiles)
}

func TestSetAndMustExist(t *testing.T) {
	cache := NewSimpleFileCache(nil)
	fileName := "test-file"
	content := []byte("test-content")

	// Test Set
	err := cache.Set(fileName, content, time.Minute)
	require.NoError(t, err)

	// Test MustExist
	entry, err := cache.MustExist(fileName)
	require.NoError(t, err)
	assert.Equal(t, fileName, entry.Name)
	assert.Equal(t, content, entry.Content)

	// Test MustExist with non-existent file
	_, err = cache.MustExist("non-existent")
	assert.Error(t, err)
}

func TestGetFile(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "filecache-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create a temporary file
	fileName := filepath.Join(tmpDir, "test.txt")
	content1 := []byte("content-1")
	err = os.WriteFile(fileName, content1, 0644)
	require.NoError(t, err)

	// Initialize cache with short freshness for disk files
	opts := &FileCacheOptions{
		FreshnessForDiskFiles: 2 * time.Second,
	}
	cache := NewSimpleFileCache(opts)

	// 1. Initial fetch
	entry, err := cache.Get(fileName)
	require.NoError(t, err)
	assert.Equal(t, content1, entry.Content)

	// Wait half a second to ensure the second write has in the file system a bigger timestamp
	time.Sleep(500 * time.Millisecond)

	// 2. Update file content
	content2 := []byte("content-2")
	err = os.WriteFile(fileName, content2, 0644)
	require.NoError(t, err)

	// 3. Fetch again immediately (should return cached content because it's fresh)
	entry, err = cache.Get(fileName)
	require.NoError(t, err)
	assert.Equal(t, content1, entry.Content)

	// 4. Wait for expiration
	time.Sleep(2 * time.Second)

	// 5. Fetch again (should return new content)
	entry, err = cache.Get(fileName)
	require.NoError(t, err)
	assert.Equal(t, content2, entry.Content)
}

func TestGetURL(t *testing.T) {
	// Mock server (use TLS because the cache only supports https://)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/file":
			if r.Header.Get("If-None-Match") == "etag-1" {
				w.WriteHeader(http.StatusNotModified)
				return
			}
			w.Header().Set("Etag", "etag-1")
			w.Write([]byte("server-content"))
		case "/error":
			w.WriteHeader(http.StatusInternalServerError)
		case "/notfound":
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Use the server's client which trusts the self-signed cert
	opts := &FileCacheOptions{
		HTTPClient: server.Client(),
	}
	cache := NewSimpleFileCache(opts)

	// 1. Initial fetch
	url := server.URL + "/file"
	entry, err := cache.Get(url)
	require.NoError(t, err)
	assert.Equal(t, []byte("server-content"), entry.Content)
	assert.Equal(t, "etag-1", entry.Etag)

	// 2. Fetch again (should use cached entry if fresh, but let's force expiration to test 304)
	// Manually expire the entry to force a re-validation
	entry.Expires = time.Now().Add(-1 * time.Hour)

	entry, err = cache.Get(url)
	require.NoError(t, err)
	assert.Equal(t, []byte("server-content"), entry.Content) // Content shouldn't change

	// 3. Test 404
	url404 := server.URL + "/notfound"
	_, err = cache.Get(url404)
	assert.Error(t, err)

	// 4. Test 500
	url500 := server.URL + "/error"
	_, err = cache.Get(url500)
	assert.Error(t, err)
}

func TestSimpleFileCacheConcurrentAccess(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "filecache-concurrent-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	fileName := filepath.Join(tmpDir, "concurrent.txt")
	err = os.WriteFile(fileName, []byte("initial-data"), 0644)
	require.NoError(t, err)

	opts := &FileCacheOptions{
		FreshnessForDiskFiles: 10 * time.Millisecond,
	}
	cache := NewSimpleFileCache(opts)

	var wg sync.WaitGroup
	const goroutines = 20
	const iterations = 50

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				entry, err := cache.Get(fileName)
				if assert.NoError(t, err) && assert.NotNil(t, entry) {
					_ = entry.Name
					_ = len(entry.Content)
					_ = entry.Expires
				}
				time.Sleep(2 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
}

func TestSingleFlightCoalescing(t *testing.T) {
	var requestCount int32
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)
		time.Sleep(50 * time.Millisecond)
		w.Header().Set("Etag", "etag-sf")
		w.Write([]byte("coalesced-content"))
	}))
	defer server.Close()

	opts := &FileCacheOptions{
		HTTPClient: server.Client(),
	}
	cache := NewSimpleFileCache(opts)
	url := server.URL + "/slow-file"

	var wg sync.WaitGroup
	const concurrentCallers = 20

	for i := 0; i < concurrentCallers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			entry, err := cache.Get(url)
			require.NoError(t, err)
			assert.Equal(t, []byte("coalesced-content"), entry.Content)
		}()
	}

	wg.Wait()
	assert.Equal(t, int32(1), atomic.LoadInt32(&requestCount), "expected singleflight to coalesce all 20 calls into 1 HTTP request")
}


