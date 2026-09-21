// Copyright 2023-2026 Jesus Ruiz. All rights reserved.
// Use of this source code is governed by an Apache 2.0
// license that can be found in the LICENSE file.

package pdp

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/DOME-Marketplace/tmfgo/internal/errl"
	"github.com/DOME-Marketplace/tmfgo/internal/filecache"

	st "go.starlark.net/starlark"
)

type Config struct {

	// PolicyFileName is the name of the file where the policies are stored.
	// It can specify a local file or a remote URL.
	PolicyFileName string

	// The http Client to retrieve the policies from a remote server.
	// If nil we use our own http.Client with a timeout of 10 seconds and no redirects.
	httpClient *http.Client

	// Debug mode, more logs and less caching
	Debug bool
}

// Validate checks if the Config is valid
func (c *Config) Validate() error {
	if c == nil {
		return errl.Errorf("config cannot be nil")
	}
	if c.PolicyFileName == "" {
		return errl.Errorf("PolicyFileName is required")
	}
	return nil
}

// PDP implements a simple Policy Decision Point in Starlark, for use in front of TMForum APIs.
//
// There can be several instances simultaneously, and each instance is safe for concurrent
// use by different goroutines.
type PDP struct {

	// The name of the file where the policy rules reside.
	scriptname string

	debug bool

	// The file cache to read the policy and other files. Modifications to the original file
	// are picked up automatically according to a freshness policy.
	fileCache filecache.FileCache

	// The pool of instances of the policy execution engines, to minimize startup
	// and teardown overheads.
	// Every goroutine uses its own instance from the pool, so they are goroutine safe.
	// If the file with the policies change, the associated Starlark thread is updated,
	// Goroutines using old versions will run until completion, and new ones will pick
	// the new version of the policies
	threadPool sync.Pool

	// Counter of thread pool entries created
	threadPoolCounter atomic.Int64

	// The http Client to retrieve the policies from a remote server if configured to do so.
	httpClient *http.Client
}

// NewPDPService creates a new PDP instance.
func NewPDPService(
	config *Config,
) (*PDP, error) {

	if err := config.Validate(); err != nil {
		return nil, errl.Errorf("invalid config: %w", err)
	}

	m := &PDP{}
	m.scriptname = config.PolicyFileName

	// Create the file cache and initialize it with the policy file.
	m.fileCache = filecache.NewSimpleFileCache(nil)
	_, err := m.fileCache.Get(config.PolicyFileName)
	if err != nil {
		return nil, errl.Errorf("error getting file cache entry for: %s: %w", config.PolicyFileName, err)
	}

	// Create the pool of parsed and compiled Starlark policy rules.
	m.threadPool = sync.Pool{
		New: func() any {
			count := m.threadPoolCounter.Add(1)
			slog.Debug("Creating a new thread entry in the PDP pool", slog.Int64("count", count))
			te, err := m.bufferedParseAndCompileFile(m.scriptname)
			if err != nil {
				slog.Error("Error creating a new thread entry in the PDP pool", slog.String("error", err.Error()))
				return nil
			}
			return te
		},
	}

	m.debug = config.Debug

	if config.httpClient != nil {
		// Use the supplied http.Client if provided.
		m.httpClient = config.httpClient
	} else {
		// We use an http.Client with a timeout of 10 seconds and no redirects.
		m.httpClient = &http.Client{
			Timeout: 10 * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return errors.New("no redirects allowed")
			},
		}
	}

	return m, nil
}

// bufferedParseAndCompileFile reads a file with Starlark code and compiles it
func (m *PDP) bufferedParseAndCompileFile(scriptname string) (*threadEntry, error) {
	threadEntry := m.createThreadEntry(scriptname)
	if threadEntry == nil {
		return nil, errl.Errorf("error creating thread entry")
	}

	entry, err := m.fileCache.Get(scriptname)
	if err != nil {
		return nil, errl.Errorf("error getting file cache entry")
	}

	threadEntry.scriptHash = entry.FileHash
	src := entry.Content

	if err := m.compileStarlarkScript(threadEntry, string(src)); err != nil {
		return nil, errl.Errorf("error compiling Starlark program")
	}

	if err := m.validateCompiledScript(threadEntry); err != nil {
		return nil, errl.Errorf("error getting authorize function")
	}

	return threadEntry, nil
}

// reset checks if the thread entry needs to be recompiled
func (m *PDP) reset(te *threadEntry) error {
	entry, err := m.fileCache.Get(te.scriptname)
	if err != nil {
		return errl.Errorf("error reading script file %s: %w", te.scriptname, err)
	}

	// If hashes are the same, we do not need to recompile the file.
	if entry.FileHash == te.scriptHash {
		return nil
	}

	// The file has changed, so we recompile it.
	src := entry.Content

	if err := m.compileStarlarkScript(te, string(src)); err != nil {
		return errl.Errorf("error compiling Starlark program: %w", err)
	}

	if err := m.validateCompiledScript(te); err != nil {
		return errl.Errorf("error getting authorize function: %w", err)
	}

	return nil
}

// Authorize evaluates authorization policies against the provided input data.
// It returns true if the request is authorized, false otherwise.
func (m *PDP) Authorize(input StarTMFMap) (bool, error) {
	return m.evaluateDecision(Authorize, input)
}

// evaluateDecision is the internal function that handles both authentication and authorization decisions
func (m *PDP) evaluateDecision(decision Decision, input StarTMFMap) (bool, error) {
	if !decision.IsValid() {
		return false, errl.Errorf("invalid decision type: %v", decision)
	}

	if input == nil {
		return false, errl.Errorf("input cannot be nil")
	}

	// Get a Starlark Thread from the pool to evaluate the policies.
	pe := m.threadPool.Get()
	if pe == nil {
		return false, errl.Errorf("getting a thread entry from pool")
	}
	defer m.threadPool.Put(pe)

	entry, ok := pe.(*threadEntry)
	if !ok {
		return false, errl.Errorf("invalid entry type in the pool")
	}

	// Check if the thread is still valid. If not, we need to recompile the file.
	err := m.reset(entry)
	if err != nil {
		return false, err
	}

	// We mutate the predeclared identifier, so the policy can access the data for this request.
	// We can also service possible callbacks from the rules engine.
	entry.predeclared["input"] = input

	// Build the arguments to the StarLark function, which is empty.
	var args st.Tuple

	// Call the corresponding function in the Starlark Thread
	var result st.Value
	if decision == Authenticate {
		// For now, we only support authorization
		return false, errl.Errorf("authentication not yet implemented")
	} else {
		// Call the 'authorize' function
		result, err = st.Call(entry.thread, entry.authorizeFunction, args, nil)
	}

	if err != nil {
		fmt.Printf("rules ERROR: %s\n", err.(*st.EvalError).Backtrace())
		return false, errl.Errorf("error calling function: %w", err)
	}

	// Check that the value returned is of the correct type (boolean)
	resultType := result.Type()
	if resultType != "bool" {
		err := errl.Errorf("function returned wrong type: %v", resultType)
		return false, err
	}

	// Return the value as a Go boolean
	return bool(result.(st.Bool).Truth()), nil
}

func (m *PDP) GetFile(filename string) (*filecache.FileEntry, error) {

	entry, err := m.fileCache.MustExist(filename)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (m *PDP) PutFile(filename string, content []byte) error {
	return m.fileCache.Set(filename, content, 0)
}
