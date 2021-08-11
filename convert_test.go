package radix

import (
	"reflect"
	"testing"
)

func TestConvertEmpty(t *testing.T) {
	out, err := Convert(nil, 16, 2)
	if out != nil {
		t.Error("out must be nil")
	}
	if err != nil {
		t.Error("err must be nil")
	}
}

func TestConvertErrors(t *testing.T) {
	out, err := Convert(nil, 1, 2)
	if out != nil {
		t.Error("out must be nil")
	}
	if err != nil {
		t.Error("err must be not nil")
	} else if err.Error() != "inrx must be greater than 1" {
		t.Error("err must be invalid inrx")
	}

	out, err = Convert(nil, 2, 1)
	if out != nil {
		t.Error("out must be nil")
	}
	if err != nil {
		t.Error("err must be not nil")
	} else if err.Error() != "outrx must be greater than 1" {
		t.Error("err must be invalid outrx")
	}

	out, err = Convert([]int{1, 2, 3}, 2, 2)
	if out != nil {
		t.Error("out must be nil")
	}
	if err != nil {
		t.Error("err must be not nil")
	} else if err.Error() != "in[1]: 2 must be less than inrx: 2" {
		t.Error("err must be digit is greater than inrx")
	}

	out, err = Convert([]int{1, -1, 3}, 2, 2)
	if out != nil {
		t.Error("out must be nil")
	}
	if err != nil {
		t.Error("err must be not nil")
	} else if err.Error() != "in[1]: -1 must be greater or equal 0" {
		t.Error("err must be digit is negative")
	}
}

func TestConvertBytesEmpty(t *testing.T) {
	out, err := ConvertBytes(nil, 58)
	if out != nil {
		t.Error("out must be nil")
	}
	if err != nil {
		t.Error("err must be nil")
	}
}

func TestConvertBytesVsConvert(t *testing.T) {
	bin := generateBytes(64, 256)
	iin := generateInts(64, 256)

	rxs := []int{2, 8, 10, 11, 29, 58, 256}

	for _, rx := range rxs {
		bout, err := ConvertBytes(bin, rx)
		if err != nil {
			t.Fatal(err)
		}
		iout, err := Convert(iin, 256, rx)
		if err != nil {
			t.Fatal(err)
		}
		if !cmpBytesVsInts(bout, iout) {
			t.Errorf("failure for radix: %v", rx)
		}
	}
}

func cmpBytesVsInts(bytes []byte, ints []int) bool {
	if len(bytes) != len(ints) {
		return false
	}
	for i := 0; i < len(bytes); i++ {
		if int(bytes[i]) != ints[i] {
			return false
		}
	}

	return true
}

func BenchmarkConvertExtraSmall(b *testing.B) {
	input := generateInts(1, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Convert(input, 256, 58)
	}
}

func BenchmarkConvertBytesExtraSmall(b *testing.B) {
	input := generateBytes(1, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConvertBytes(input, 58)
	}
}

func BenchmarkConvertSmall(b *testing.B) {
	input := generateInts(20, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Convert(input, 256, 58)
	}
}

func BenchmarkConvertBytesSmall(b *testing.B) {
	input := generateBytes(20, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConvertBytes(input, 58)
	}
}

func BenchmarkConvertMedium(b *testing.B) {
	input := generateInts(1024, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Convert(input, 256, 58)
	}
}

func BenchmarkConvertBytesMedium(b *testing.B) {
	input := generateBytes(1024, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConvertBytes(input, 58)
	}
}

func BenchmarkConvertLarge(b *testing.B) {
	input := generateInts(16*1024, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Convert(input, 256, 58)
	}
}

func BenchmarkConvertBytesLarge(b *testing.B) {
	input := generateBytes(16*1024, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConvertBytes(input, 58)
	}
}

func generateInts(n int, rx int) (out []int) {
	out = make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = (rx - 1) - i%rx
	}
	return
}

func generateBytes(n int, rx int) (out []byte) {
	out = make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = byte((rx - 1) - i%rx)
	}
	return
}

func TestGenerateInts(t *testing.T) {
	tests := []struct {
		len   int
		radix int
		want  []int
	}{
		{len: 1, radix: 10, want: []int{9}},
		{len: 4, radix: 1000, want: []int{999, 998, 997, 996}},
		{len: 10, radix: 5, want: []int{4, 3, 2, 1, 0, 4, 3, 2, 1, 0}},
	}

	for _, test := range tests {
		res := generateInts(test.len, test.radix)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("generateInts(1, 10): %v not equal %v", res, test.want)
		}
	}
}

func TestGenerateBytes(t *testing.T) {
	tests := []struct {
		len   int
		radix int
		want  []byte
	}{
		{len: 1, radix: 10, want: []byte{9}},
		{len: 4, radix: 256, want: []byte{255, 254, 253, 252}},
		{len: 10, radix: 5, want: []byte{4, 3, 2, 1, 0, 4, 3, 2, 1, 0}},
	}

	for _, test := range tests {
		res := generateBytes(test.len, test.radix)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("generateInts(1, 10): %v not equal %v", res, test.want)
		}
	}
}
