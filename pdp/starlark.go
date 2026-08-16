// Copyright 2023-2025 Jesus Ruiz. All rights reserved.
// Use of this source code is governed by an Apache 2.0
// license that can be found in the LICENSE file.

package pdp

import (
	"io"
	"log"
	"log/slog"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/hesusruiz/tmforum/internal/errl"
	starjson "go.starlark.net/lib/json"
	"go.starlark.net/lib/math"
	sttime "go.starlark.net/lib/time"
	"go.starlark.net/repl"
	st "go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
	"go.starlark.net/syntax"
)

func init() {
	// Add our built-ins to the Starlark Universe dictionary before any evaluation begins.
	// All values here must be immutable and shared among all instances.
	// See here for the standard Starlark entities:
	// https://github.com/google/starlark-go/blob/master/doc/spec.md#built-in-constants-and-functions

	// Create a StarLark module with our own utility functions
	var Module = &starlarkstruct.Module{
		Name: "star",
		Members: st.StringDict{
			"getinput": st.NewBuiltin("getinput", getInputElement),
			"getbody":  st.NewBuiltin("getbody", getRequestBody),
		},
	}

	// Set the global Starlark environment with required modules, including our own
	st.Universe["json"] = starjson.Module
	st.Universe["time"] = sttime.Module
	st.Universe["math"] = math.Module
	st.Universe["star"] = Module
}

// threadEntry is one entry in a pool of Starlark threads for policy rules execution.
// All instances are normally the same, using the same compiled version of the same file.
// The pool increases concurrency because a given Starlark thread can be reused
// among goroutines, but not used concurrently by the same goroutine.
// Another benefit is that it facilitates the dynamic update of policy files without
// affecting concurrency.
type threadEntry struct {
	globals           st.StringDict
	predeclared       st.StringDict
	thread            *st.Thread
	authorizeFunction *st.Function
	scriptname        string
	scriptHash        uint64
}

// createThreadEntry creates a new thread entry with basic initialization
func (m *PDP) createThreadEntry(scriptname string) *threadEntry {
	te := &threadEntry{}
	te.scriptname = scriptname

	logger := slog.Default()

	// The compiled program context will be stored in a new Starlark thread for each invocation
	te.thread = &st.Thread{
		Load: repl.MakeLoadOptions(&syntax.FileOptions{}),
		Print: func(_ *st.Thread, msg string) {
			logger.Info("rules => " + msg)
		},
		Name: "exec " + scriptname,
	}

	// Create a predeclared environment holding the 'input' object.
	// For the moment it is empty, but it will be mutated for each request for authentication.
	te.predeclared = st.StringDict{}
	te.predeclared["input"] = StarTMFMap{}

	return te
}

// compileStarlarkScript compiles the Starlark source code
func (m *PDP) compileStarlarkScript(te *threadEntry, src string) error {
	// Parse and execute the top-level commands in the script file
	// The globals are thread-local and not process-global
	var err error
	te.globals, err = st.ExecFileOptions(&syntax.FileOptions{}, te.thread, te.scriptname, src, te.predeclared)
	if err != nil {
		return errl.Errorf("error compiling Starlark program: %w", err)
	}

	// Make sure that the global environment is frozen so the Startlark script cannot
	// modify it. This is important for security and to avoid concurrency problems.
	te.globals.Freeze()

	return nil
}

// validateCompiledScript validates that the compiled script has the required functions
func (m *PDP) validateCompiledScript(te *threadEntry) error {
	// The module has to define a function called 'authorize', which will be invoked
	// for each request to access protected resources.
	var err error
	te.authorizeFunction, err = getGlobalFunction(te.globals, "authorize")
	if err != nil {
		return errl.Errorf("error getting authorize function: %w", err)
	}

	return nil
}

// getGlobalFunction retrieves a Callable from the supplied globals dictionary.
func getGlobalFunction(globals st.StringDict, funcName string) (*st.Function, error) {

	// Check that we have the function
	f, ok := globals[funcName]
	if !ok {
		err := errl.Errorf("missing definition of %s", funcName)
		log.Println(err.Error())
		return nil, err
	}

	// Check that is is a Callable
	starFunction, ok := f.(*st.Function)
	if !ok {
		err := errl.Errorf("expected a Callable but got %v", f.Type())
		log.Println(err.Error())
		return nil, err
	}

	return starFunction, nil
}

// getInputElement is a Starlark builtin function to get input elements
func getInputElement(thread *st.Thread, _ *st.Builtin, args st.Tuple, kwargs []st.Tuple) (st.Value, error) {

	// Get the current input structure being processed
	r := thread.Local("inputrequest")
	input, ok := r.(StarTMFMap)
	if !ok {
		return st.None, errl.Errorf("no request found in thread locals")
	}

	// Get the element
	var elemPath string
	err := st.UnpackPositionalArgs("input2", args, kwargs, 1, &elemPath)
	if err != nil {
		return nil, err
	}

	elem, err := GetValue(input, elemPath)
	if err != nil {
		return st.None, nil
	}
	return elem, nil
}

// getRequestBody is a Starlark builtin function to get request body
func getRequestBody(thread *st.Thread, _ *st.Builtin, args st.Tuple, kwargs []st.Tuple) (st.Value, error) {

	// Get the current HTTP request being processed
	r := thread.Local("httprequest")
	request, ok := r.(*http.Request)
	if !ok {
		return st.None, errl.Errorf("no request found in thread locals")
	}

	// Read the body from the request and store in thread locals in case we need it later
	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, errl.Errorf("error reading request body: %w", err)
	}
	thread.SetLocal("requestbody", bytes)

	// Return string for the Starlark script
	body := st.String(bytes)

	return body, nil
}

// Get returns a child of the given value according to a dotted path.
// The source data must be either map[string]any or []any
func GetValue(a StarTMFMap, path string) (st.Value, error) {
	if a == nil {
		return st.None, errl.Errorf("input map cannot be nil")
	}

	if path == "" {
		return st.None, errl.Errorf("path cannot be empty")
	}

	parts := strings.Split(path, ".")
	var src st.Value = a

	// Get the value.
	for pos, pathComponent := range parts {
		var err error
		src, err = getValueAtPath(src, pathComponent, parts[:pos+1])
		if err != nil {
			return nil, err
		}
		if src == st.None {
			return st.None, nil
		}
	}

	return src, nil
}

// getValueAtPath retrieves a value at a specific path component
func getValueAtPath(src st.Value, pathComponent string, pathSoFar []string) (st.Value, error) {
	switch src.Type() {
	case "tmfmap":
		return getValueFromMap(src.(StarTMFMap), pathComponent)
	case "tmflist":
		return getValueFromList(src.(StarTMFList), pathComponent, pathSoFar)
	default:
		return nil, errl.Errorf(
			"jpath.Get: invalid type at %q: expected []any or map[string]any; got %T",
			strings.Join(pathSoFar, "."), src)
	}
}

// getValueFromMap retrieves a value from a map
func getValueFromMap(m StarTMFMap, key string) (st.Value, error) {
	if value, ok := m[key]; ok {
		return anyToValue(value), nil
	}
	return st.None, nil
}

// getValueFromList retrieves a value from a list by index
func getValueFromList(l StarTMFList, pathComponent string, pathSoFar []string) (st.Value, error) {
	// If data is an array, the path component must be an integer (base 10) to index the array
	index, err := strconv.ParseInt(pathComponent, 10, 0)
	if err != nil {
		return nil, errl.Errorf("jpath.Get: invalid list index at %q",
			strings.Join(pathSoFar, "."))
	}
	if int(index) < len(l) {
		// Update src to be the indexed element of the array
		value := l[index]
		return anyToValue(value), nil
	} else {
		return nil, errl.Errorf(
			"jpath.Get: index out of range at %q: list has only %v items",
			strings.Join(pathSoFar, "."), len(l))
	}
}

// anyToValue converts a Go value to a Starlark value
func anyToValue(value any) st.Value {

	// Special case for strings
	vvvv := reflect.ValueOf(value)
	if vvvv.Kind() == reflect.String {
		return st.String(vvvv.String())
	}

	switch v := value.(type) {
	case StarTMFMap:
		return StarTMFMap(v)
	case StarTMFList:
		return StarTMFList(v)
	case string:
		return st.String(v)
	case st.String:
		return st.String(v)
	case map[string]any:
		return StarTMFMap(v)
	case []any:
		var l []st.Value
		for _, elem := range v {
			l = append(l, anyToValue(elem))
		}
		return StarTMFList(l)
	case bool:
		return st.Bool(v)
	case float64:
		return st.Float(v)
	case int:
		return st.MakeInt(v)
	default:
		return st.None
	}
}
