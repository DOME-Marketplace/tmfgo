package errl

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

func TestError(t *testing.T) {
	err := errors.New("test error")
	e := Error(err)

	if e == nil {
		t.Fatal("expected non-nil error")
	}

	if e.Unwrap() != err {
		t.Errorf("expected underlying error %v, got %v", err, e.Unwrap())
	}

	if e.Naked() != err {
		t.Errorf("expected naked error %v, got %v", err, e.Naked())
	}

	errMsg := e.Error()
	if !strings.Contains(errMsg, "errl_test.go") {
		t.Errorf("expected error message to contain file name, got: %s", errMsg)
	}
	if !strings.Contains(errMsg, "TestError") {
		t.Errorf("expected error message to contain function name, got: %s", errMsg)
	}
	if !strings.Contains(errMsg, "test error") {
		t.Errorf("expected error message to contain original error message, got: %s", errMsg)
	}

	if nilErr := Error(nil); nilErr != nil {
		t.Errorf("expected nil for nil error, got %v", nilErr)
	}
}

func TestErrorf(t *testing.T) {
	e := Errorf("formatted error: %d", 123)

	if e == nil {
		t.Fatal("expected non-nil error")
	}

	errMsg := e.Error()
	if !strings.Contains(errMsg, "errl_test.go") {
		t.Errorf("expected error message to contain file name, got: %s", errMsg)
	}
	if !strings.Contains(errMsg, "TestErrorf") {
		t.Errorf("expected error message to contain function name, got: %s", errMsg)
	}
	if !strings.Contains(errMsg, "formatted error: 123") {
		t.Errorf("expected error message to contain formatted message, got: %s", errMsg)
	}
}

func helperError2(err error) *ErrorWithLocation {
	return Error2(err)
}

func helperErrorf2(format string, a ...any) *ErrorWithLocation {
	return Errorf2(format, a...)
}

func TestError2(t *testing.T) {
	err := errors.New("test error 2")
	e := helperError2(err)

	if e == nil {
		t.Fatal("expected non-nil error")
	}

	errMsg := e.Error()
	// Should point to TestError2, not helperError2
	if !strings.Contains(errMsg, "TestError2") {
		t.Errorf("expected error message to contain calling function name (TestError2), got: %s", errMsg)
	}

	if nilErr := Error2(nil); nilErr != nil {
		t.Errorf("expected nil for nil error, got %v", nilErr)
	}
}

func TestErrorf2(t *testing.T) {
	e := helperErrorf2("formatted error 2: %s", "hello")

	if e == nil {
		t.Fatal("expected non-nil error")
	}

	errMsg := e.Error()
	// Should point to TestErrorf2, not helperErrorf2
	if !strings.Contains(errMsg, "TestErrorf2") {
		t.Errorf("expected error message to contain calling function name (TestErrorf2), got: %s", errMsg)
	}
}

func TestSeverityLevel(t *testing.T) {
	tests := []struct {
		severity SeverityLevel
		expected string
	}{
		{DebugM, "Debug"},
		{InfoM, "Info"},
		{WarnM, "Warn"},
		{ErrorM, "Error"},
		{SeverityLevel(99), "Unknown"},
	}

	for _, tc := range tests {
		if tc.severity.String() != tc.expected {
			t.Errorf("expected %s for severity %d, got %s", tc.expected, tc.severity, tc.severity.String())
		}
	}
}

func TestValidationMessages(t *testing.T) {
	var msgs ValidationMessages
	msgs.Add(InfoM, "message 1")
	msgs.Addf(ErrorM, "message %d", 2)

	if len(msgs) != 2 {
		t.Errorf("expected 2 messages, got %d", len(msgs))
	}

	expectedString := "Info: message 1\nError: message 2\n"
	if msgs.String() != expectedString {
		t.Errorf("expected string:\n%s\ngot:\n%s", expectedString, msgs.String())
	}
}

func TestJsonUnmarshalError(t *testing.T) {
	data := []byte(`{"name": "test", "age": "not-a-number"}`)
	type Target struct {
		Age int `json:"age"`
	}
	var target Target
	err := json.Unmarshal(data, &target)

	if err == nil {
		t.Fatal("expected unmarshal error")
	}

	e := JsonUnmarshalError(data, err)
	if e == nil {
		t.Fatal("expected non-nil result from JsonUnmarshalError")
	}

	errMsg := e.Error()
	if !strings.Contains(errMsg, "offset") {
		t.Errorf("expected error message to contain 'offset', got: %s", errMsg)
	}

	// Test syntax error in the middle of data
	syntaxData := []byte(`{"name": "test",, "age": 1}`) // double comma
	err = json.Unmarshal(syntaxData, &target)
	if err == nil {
		t.Fatal("expected syntax error")
	}
	e = JsonUnmarshalError(syntaxData, err)
	if !strings.Contains(e.Error(), "offset") {
		t.Errorf("expected error message to contain 'offset' for syntax error, got: %s", e.Error())
	}

	// Test syntax error at the beginning (start < 0)
	syntaxDataSmall := []byte(`{,}`) // error at offset 1
	err = json.Unmarshal(syntaxDataSmall, &target)
	if err == nil {
		t.Fatal("expected syntax error")
	}
	e = JsonUnmarshalError(syntaxDataSmall, err)
	if !strings.Contains(e.Error(), "offset") {
		t.Errorf("expected error message to contain 'offset' for small syntax error, got: %s", e.Error())
	}

	// Test UnmarshalTypeError at the beginning (start < 0)
	typeDataSmall := []byte(`{"a":"b"}`)
	type TargetSmall struct {
		A int `json:"a"`
	}
	var targetSmall TargetSmall
	err = json.Unmarshal(typeDataSmall, &targetSmall)
	if err == nil {
		t.Fatal("expected type error")
	}
	e = JsonUnmarshalError(typeDataSmall, err)
	if !strings.Contains(e.Error(), "offset") {
		t.Errorf("expected error message to contain 'offset' for small type error, got: %s", e.Error())
	}
	if strings.Contains(e.Error(), "...") {
		t.Errorf("expected no ellipsis for small type error, got: %s", e.Error())
	}

	// Test case where err is not SyntaxError or UnmarshalTypeError
	e = JsonUnmarshalError([]byte(`{}`), errors.New("some other error"))
	if e == nil {
		t.Fatal("expected non-nil result from JsonUnmarshalError for other error")
	}

	// Test offset at end of data (specifically for the if offset >= len branches)
	e = JsonUnmarshalError([]byte("{}"), &json.SyntaxError{Offset: 10})
	if e == nil {
		t.Error("expected non-nil error")
	}
	e = JsonUnmarshalError([]byte("{}"), &json.UnmarshalTypeError{Offset: 10})
	if e == nil {
		t.Error("expected non-nil error")
	}
}
