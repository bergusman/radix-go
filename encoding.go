package radix

import (
	"errors"
	"fmt"
)

func Encode(input []int, alphabet string) (output string, err error) {
	runes := []rune(alphabet)
	if len(runes) < 2 {
		return "", errors.New("len(alphabet) less then 2")
	}
	if len(input) == 0 {
		return "", nil
	}

	for i, v := range input {
		if v >= len(runes) {
			return "", fmt.Errorf("input[%v]: %v must be less then len(alphabet): %v", i, v, len(runes))
		}
		output += string(runes[v])
	}
	return
}

func Decode(input string, alphabet string) (output []int, err error) {
	runes := []rune(alphabet)
	if len(runes) < 2 {
		return nil, errors.New("len(alphabet) less then 2")
	}

	if input == "" {
		return nil, nil
	}

	runesMap := make(map[rune]int)
	for i, r := range runes {
		runesMap[r] = i
	}

	for i, r := range input {
		if v, ok := runesMap[r]; ok {
			output = append(output, v)
		} else {
			return nil, fmt.Errorf("rune %q at %v not contained in alphabet", r, i)
		}
	}

	return
}

func EncodeBytes(input []byte, alphabet string) (output string, err error) {
	runes := []rune(alphabet)
	if len(runes) < 2 {
		return "", errors.New("len(alphabet) less then 2")
	}
	if len(input) == 0 {
		return "", nil
	}

	for i, v := range input {
		if int(v) >= len(runes) {
			return "", fmt.Errorf("input[%v]: %v must be less then len(alphabet): %v", i, v, len(runes))
		}
		output += string(runes[v])
	}
	return
}

func DecodeBytes(input string, alphabet string) (output []byte, err error) {
	runes := []rune(alphabet)
	if len(runes) < 2 {
		return nil, errors.New("len(alphabet) less then 2")
	}

	if input == "" {
		return nil, nil
	}

	runesMap := make(map[rune]byte)
	for i, r := range runes {
		if i > 255 {
			break
		}
		runesMap[r] = byte(i)
	}

	for i, r := range input {
		if v, ok := runesMap[r]; ok {
			output = append(output, v)
		} else {
			return nil, fmt.Errorf("rune %q at %v not contained in alphabet", r, i)
		}
	}

	return
}
