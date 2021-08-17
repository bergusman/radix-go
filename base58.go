package radix

import (
	"errors"
	"unicode/utf8"
)

var ErrBase58BadAlphabet = errors.New("base58: alphabet length less than 58")

const AlphabetBitcoin = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const AlphabetRipple = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
const AlphabetFlickr = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

func Base58Encode(input []byte, alphabet string) (string, error) {
	if utf8.RuneCountInString(alphabet) < 58 {
		return "", ErrBase58BadAlphabet
	}

	input58, err := ConvertBytes(input, 58)
	if err != nil {
		return "", err
	}

	zeror, _ := utf8.DecodeRuneInString(alphabet)
	zero := string(zeror)
	padding := ""
	for _, in := range input {
		if in == 0 {
			padding += zero
		} else {
			break
		}
	}

	output, err := EncodeBytes(input58, alphabet)
	if err != nil {
		return "", err
	}
	if len(padding) == 0 {
		return output, nil
	}
	return padding + output, nil
}

func Base58Decode(input string, alphabet string) ([]byte, error) {
	if utf8.RuneCountInString(alphabet) < 58 {
		return nil, ErrBase58BadAlphabet
	}
	return DecodeBytes(input, alphabet)
}
