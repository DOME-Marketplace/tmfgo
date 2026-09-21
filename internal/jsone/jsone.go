package jsone

import (
	"encoding/json"
	"fmt"

	"github.com/DOME-Marketplace/tmfgo/internal/errl"
)

func JsonUnmarshalError(data []byte, err error) error {

	if syntaxErr, ok := err.(*json.SyntaxError); ok {
		offset := syntaxErr.Offset
		if offset >= int64(len(data)) {
			return errl.Error2(err)
		}
		start := offset - 20
		if start < 0 {
			return errl.Errorf2("%s<--[offset:%d] %s", data[:offset], offset, err)
		}
		return errl.Errorf2("... %s<--[offset:%d] %s", data[start:offset], offset, err)
	} else if typeErr, ok := err.(*json.UnmarshalTypeError); ok {
		offset := typeErr.Offset
		if offset >= int64(len(data)) {
			return errl.Error2(err)
		}
		start := offset - 20
		if start < 0 {
			return errl.Errorf2("%s<--[offset:%d] %s", data[:offset], offset, err)
		}
		return errl.Errorf2("... %s<--[offset:%d] %s", data[start:offset], offset, err)
	} else {
		return errl.Error2(err)
	}

}

// Unmarshal is a wrapper around json.Unmarshal that adds context to the error message.
func Unmarshal(data []byte, v any) error {

	err := json.Unmarshal(data, v)
	if err == nil {
		return nil
	}

	var offset, start int64

	if syntaxErr, ok := err.(*json.SyntaxError); ok {
		offset = syntaxErr.Offset
	} else if typeErr, ok := err.(*json.UnmarshalTypeError); ok {
		offset = typeErr.Offset
	} else {
		return err
	}

	start = offset - 20
	if start < 0 {
		return fmt.Errorf("%s<--[offset:%d] %s", data[:offset], offset, err)
	}
	return fmt.Errorf("... %s<--[offset:%d] %s", data[start:offset], offset, err)

}
