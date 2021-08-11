package radix

import (
	"testing"
)

func TestEncodeEmpty(t *testing.T) {
	out, err := Encode(nil, "01")
	if out != "" {
		t.Error("out must be empty string")
	}
	if err != nil {
		t.Error("err must be nil")
	}
}

func TestEncodeErrors(t *testing.T) {
	out, err := Encode(nil, "0")
	if out != "" {
		t.Error("out must be empty string")
	}
	if err == nil {
		t.Error("err must be not nil")
	} else if err.Error() != "len(alphabet) less than 2" {
		t.Error("len(alphabet) less than 2")
	}
}
