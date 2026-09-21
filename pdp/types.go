// Copyright 2023 Jesus Ruiz. All rights reserved.
// Use of this source code is governed by an Apache 2.0
// license that can be found in the LICENSE file.

package pdp

import (
	"fmt"
	"hash/maphash"
	"strings"

	"github.com/DOME-Marketplace/tmfgo/internal/errl"
	st "go.starlark.net/starlark"
)

var seed = maphash.MakeSeed()

// hashString computes the hash of s.
func hashString(s string) uint32 {
	if len(s) >= 12 {
		// Call the Go runtime's optimized hash implementation,
		// which uses the AES instructions on amd64 and arm64 machines.
		h := maphash.String(seed, s)
		return uint32(h>>32) | uint32(h)
	}
	return softHashString(s)
}

// softHashString computes the 32-bit FNV-1a hash of s in software.
func softHashString(s string) uint32 {
	var h uint32 = 2166136261
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

// Decision represents the type of decision to be made
type Decision int

const (
	// Authenticate represents an authentication decision
	Authenticate Decision = 1
	// Authorize represents an authorization decision
	Authorize Decision = 2
)

// String returns a string representation of the Decision
func (d Decision) String() string {
	switch d {
	case Authenticate:
		return "Authenticate"
	case Authorize:
		return "Authorize"
	default:
		return "Unknown"
	}
}

// IsValid checks if the Decision value is valid
func (d Decision) IsValid() bool {
	return d == Authenticate || d == Authorize
}

// StarTMFMap represents a TMForum map that can be used in Starlark scripts
type StarTMFMap map[string]any

// Value interface
func (s StarTMFMap) String() string {
	out := new(strings.Builder)

	out.WriteByte('{')
	sep := ""
	for k, v := range s {
		out.WriteString(sep)
		s := strings.ReplaceAll(fmt.Sprintf("%v", k), " ", "")
		out.WriteString(s)
		out.WriteString(": ")

		val := anyToValue(v)
		s = strings.ReplaceAll(fmt.Sprintf("%v", val.String()), " ", "")
		out.WriteString(s)
		sep = ", "
	}
	out.WriteByte('}')
	return out.String()

}
func (s StarTMFMap) GoString() string      { return s["id"].(string) }
func (s StarTMFMap) Type() string          { return "tmfmap" }
func (s StarTMFMap) Freeze()               {} // immutable
func (s StarTMFMap) Truth() st.Bool        { return len(s) > 0 }
func (s StarTMFMap) Hash() (uint32, error) { return hashString(s["id"].(string)), nil }

// Indexable interface
func (s StarTMFMap) Len() int { return len(s) } // number of entries

// Mapping interface
func (s StarTMFMap) Get(name st.Value) (v st.Value, found bool, err error) {

	path := string(name.(st.String))

	// We need at least one name
	if path == "" {
		return s, false, nil
	}

	// This is a special case, where we assume the meaning of "this object".
	if path == "." {
		return s, true, nil
	}

	// Two consecutive dots is an error
	if strings.Contains(path, "..") {
		return nil, false, errl.Errorf("invalid path %q: contains '..'", path)
	}

	vv, err := GetValue(s, string(name.(st.String)))
	if err != nil {
		return nil, false, err
	}
	v = anyToValue(vv)
	return v, true, nil
}

// HasAttrs interface
func (s StarTMFMap) Attr(name string) (st.Value, error) {
	value, ok := s[name]
	if !ok {
		return nil, nil
	}

	return anyToValue(value), nil

}

func (s StarTMFMap) AttrNames() []string {
	var keys []string
	for key := range s {
		keys = append(keys, key)
	}
	return keys
}

// StarTMFList represents a TMForum list that can be used in Starlark scripts
type StarTMFList []st.Value

// Value interface
func (s StarTMFList) String() string {
	out := new(strings.Builder)

	out.WriteByte('[')
	for i, elem := range s {
		if i > 0 {
			out.WriteString(", ")
		}
		s := strings.ReplaceAll(fmt.Sprintf("%v", elem), " ", "")
		out.WriteString(s)
	}
	out.WriteByte(']')

	return out.String()
}
func (s StarTMFList) Type() string          { return "tmflist" }
func (s StarTMFList) Freeze()               {} // immutable
func (s StarTMFList) Truth() st.Bool        { return len(s) > 0 }
func (s StarTMFList) Hash() (uint32, error) { return hashString("tmflist"), nil }

// Indexable interface
func (s StarTMFList) Len() int { return len(s) } // number of entries
func (s StarTMFList) Index(i int) st.Value {
	value := s[i]
	return anyToValue(value)
}
