package radix

import "errors"

func Encode(input []int, alphabet string) (output string, err error) {
	if len(input) == 0 {
		return "", nil
	}
	runes := []rune(alphabet)
	if len(runes) < 2 {
		return "", errors.New("len(alphabet) less then 2")
	}
	for _, v := range input {
		output += string(runes[v])
	}
	return
}

func Decode(input []int, alphabet string) (output []int, err error) {
	return
}

func EncodeBytes(input []byte, alphabet string) (output string, err error) {
	if len(input) == 0 {
		return "", nil
	}
	runes := []rune(alphabet)
	if len(runes) < 2 {
		return "", errors.New("len(alphabet) less then 2")
	}
	if len(runes) > 256 {
		return "", errors.New("len(alphabet) greater then 256")
	}
	for _, v := range input {
		output += string(runes[v])
	}
	return
}

func DecodeBytes(input []byte, alphabet string) (output []byte, err error) {
	return
}
