package radix

import (
	"reflect"
	"testing"
)

var base58Tests = []struct {
	Bytes  []byte
	String string
}{
	{Bytes: []byte{255, 255, 255, 255}, String: "7YXq9G"},
	{Bytes: []byte{0, 0, 0, 0, 255, 255, 255, 255}, String: "11117YXq9G"},
	{Bytes: []byte{0, 1, 2, 3, 4}, String: "12VfUX"},
}

func TestBase58Encode(t *testing.T) {
	for _, tt := range base58Tests {
		output, err := Base58Encode(tt.Bytes, AlphabetBitcoin)
		if err != nil {
			t.Fatal(err)
		}
		if output != tt.String {
			t.Errorf("got: %v; want: %v", output, tt.String)
		}
	}
}

func TestBase58EncodeErrors(t *testing.T) {
	_, err := Base58Encode([]byte{}, "1234")
	if err != nil {
		if err != ErrBase58BadAlphabet {
			t.Errorf("got: %v; want: ErrBase58BadAlphabet", err)
		}
	} else {
		t.Error("want not nit err")
	}
}

func TestBase58Decode(t *testing.T) {
	for _, tt := range base58Tests {
		output, err := Base58Decode(tt.String, AlphabetBitcoin)
		if err != nil {
			t.Fatal(err)
		}
		if reflect.DeepEqual(output, tt.Bytes) {
			t.Errorf("got: %v; want: %v", output, tt.Bytes)
		}
	}
}

func TestBase58DecodeErrors(t *testing.T) {
	_, err := Base58Decode("", "1234")
	if err != nil {
		if err != ErrBase58BadAlphabet {
			t.Errorf("got: %v; want: ErrBase58BadAlphabet", err)
		}
	} else {
		t.Error("want not nit err")
	}
}
