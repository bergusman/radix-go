package radix

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func ExampleEncode() {
	out, err := Convert([]int{1, 3, 3, 7}, 10, 2)
	if err != nil {
		log.Fatal(err)
	}

	str, err := Encode(out, "01")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str)
	// Output: 10100111001
}

type encodingTest struct {
	decoded  []int
	encoded  string
	alphabet string
}

var encodingTests = []encodingTest{
	{
		decoded:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		encoded:  "1234567890",
		alphabet: "0123456789",
	},
	{
		decoded:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		encoded:  "1234567890",
		alphabet: "0123456789ABCDEF",
	},
	{
		decoded:  []int{1, 3, 3, 0, 4},
		encoded:  "🌘🌖🌖🌑🌕",
		alphabet: "🌑🌘🌗🌖🌕",
	},
}

type encodingBytesTest struct {
	decoded  []byte
	encoded  string
	alphabet string
}

var encodingBytesTests = []encodingBytesTest{
	{
		decoded:  []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		encoded:  "1234567890",
		alphabet: "0123456789",
	},
	{
		decoded:  []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		encoded:  "1234567890",
		alphabet: "0123456789ABCDEF",
	},
	{
		decoded:  []byte{1, 3, 3, 0, 4},
		encoded:  "🌘🌖🌖🌑🌕",
		alphabet: "🌑🌘🌗🌖🌕",
	},
}

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

	out, err = Encode([]int{1, 10}, "01")
	if out != "" {
		t.Error("out must be empty string")
	}
	if err == nil {
		t.Error("err must be not nil")
	} else if err.Error() != "input[1]: 10 must be less than len(alphabet): 2" {
		t.Error("err must be bad input and alphabet")
	}
}

func TestEncode(t *testing.T) {
	for _, test := range encodingTests {
		out, err := Encode(test.decoded, test.alphabet)
		if err != nil {
			t.Error(out)
		}
		if out != test.encoded {
			t.Errorf("%v != %v", out, test.decoded)
		}
	}
}

func TestEncodeBytesEmpty(t *testing.T) {
	out, err := EncodeBytes(nil, "01")
	if out != "" {
		t.Error("out must be empty string")
	}
	if err != nil {
		t.Error("err must be nil")
	}
}

func TestEncodeBytesErrors(t *testing.T) {
	out, err := EncodeBytes(nil, "0")
	if out != "" {
		t.Error("out must be empty string")
	}
	if err == nil {
		t.Error("err must be not nil")
	} else if err.Error() != "len(alphabet) less than 2" {
		t.Error("len(alphabet) less than 2")
	}

	out, err = EncodeBytes([]byte{1, 10}, "01")
	if out != "" {
		t.Error("out must be empty string")
	}
	if err == nil {
		t.Error("err must be not nil")
	} else if err.Error() != "input[1]: 10 must be less than len(alphabet): 2" {
		t.Error("err must be bad input and alphabet")
	}
}

func TestEncodeBytes(t *testing.T) {
	for _, test := range encodingBytesTests {
		out, err := EncodeBytes(test.decoded, test.alphabet)
		if err != nil {
			t.Error(out)
		}
		if out != test.encoded {
			t.Errorf("%v != %v", out, test.decoded)
		}
	}
}

func TestDecodeEmpty(t *testing.T) {
	out, err := Decode("", "01")
	if out != nil {
		t.Error("out must be nil")
	}
	if err != nil {
		t.Error("err must be nil")
	}
}

func TestDecodeErrors(t *testing.T) {
	out, err := Decode("", "0")
	if out != nil {
		t.Error("out must be nil")
	}
	if err == nil {
		t.Error("err must be not nil")
	} else if err.Error() != "len(alphabet) less than 2" {
		t.Error("len(alphabet) less than 2")
	}

	out, err = Decode("1x", "01")
	if out != nil {
		t.Error("out must be nil")
	}
	if err == nil {
		t.Error("err must be not nil")
	} else if err.Error() != "rune 'x' at 1 not contained in alphabet" {
		t.Error("err must be input rune missed in alphabet")
	}
}

func TestDecode(t *testing.T) {
	for _, test := range encodingTests {
		out, err := Decode(test.encoded, test.alphabet)
		if err != nil {
			t.Error(out)
		}
		if !reflect.DeepEqual(out, test.decoded) {
			t.Errorf("%v != %v", out, test.decoded)
		}
	}
}

func TestDecodeBytesEmpty(t *testing.T) {
	out, err := DecodeBytes("", "01")
	if out != nil {
		t.Error("out must be nil")
	}
	if err != nil {
		t.Error("err must be nil")
	}
}

func TestDecodeBytesErrors(t *testing.T) {
	out, err := DecodeBytes("", "0")
	if out != nil {
		t.Error("out must be nil")
	}
	if err == nil {
		t.Error("err must be not nil")
	} else if err.Error() != "len(alphabet) less than 2" {
		t.Error("len(alphabet) less than 2")
	}

	out, err = DecodeBytes("1x", "01")
	if out != nil {
		t.Error("out must be nil")
	}
	if err == nil {
		t.Error("err must be not nil")
	} else if err.Error() != "rune 'x' at 1 not contained in alphabet" {
		t.Error("err must be input rune missed in alphabet")
	}
}

func TestDecodeBytes(t *testing.T) {
	for _, test := range encodingBytesTests {
		out, err := DecodeBytes(test.encoded, test.alphabet)
		if err != nil {
			t.Error(out)
		}
		if !reflect.DeepEqual(out, test.decoded) {
			t.Errorf("%v != %v", out, test.decoded)
		}
	}
}
