package radix

import "testing"

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
	if err.Error() != "inrx must be greater then 1" {
		t.Error("err must be invalid inrx")
	}

	out, err = Convert(nil, 2, 1)
	if out != nil {
		t.Error("out must be nil")
	}
	if err.Error() != "outrx must be greater then 1" {
		t.Error("err must be invalid outrx")
	}

	out, err = Convert([]int{1, 2, 3}, 2, 2)
	if out != nil {
		t.Error("out must be nil")
	}
	if err.Error() != "in[1]: 2 must be less then inrx: 2" {
		t.Error("err must be digit is greater then inrx")
	}

	out, err = Convert([]int{1, -1, 3}, 2, 2)
	if out != nil {
		t.Error("out must be nil")
	}
	if err.Error() != "in[1]: -1 must be greater or equal 0" {
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
	bin := benchBytes(64)
	iin := benchInts(64)

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
	input := benchInts(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Convert(input, 256, 58)
	}
}

func BenchmarkConvertBytesExtraSmall(b *testing.B) {
	input := benchBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConvertBytes(input, 58)
	}
}

func BenchmarkConvertSmall(b *testing.B) {
	input := benchInts(20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Convert(input, 256, 58)
	}
}

func BenchmarkConvertBytesSmall(b *testing.B) {
	input := benchBytes(20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConvertBytes(input, 58)
	}
}

func BenchmarkConvertMedium(b *testing.B) {
	input := benchInts(1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Convert(input, 256, 58)
	}
}

func BenchmarkConvertBytesMedium(b *testing.B) {
	input := benchBytes(1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConvertBytes(input, 58)
	}
}

func BenchmarkConvertLarge(b *testing.B) {
	input := benchInts(16 * 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Convert(input, 256, 58)
	}
}

func BenchmarkConvertBytesLarge(b *testing.B) {
	input := benchBytes(16 * 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConvertBytes(input, 58)
	}
}

func benchInts(n int) (out []int) {
	out = make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = 255 - i%256
	}
	return
}

func benchBytes(n int) (out []byte) {
	out = make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = byte(255 - i%256)
	}
	return
}