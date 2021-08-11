// Package radix implements radix conversions for integer sequences ([]byte or []int)
// and their encoding/deconding to string by specified alphabet.
package radix

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

// Convert converts input (int) with radix (inrx) to ouput (out) with radix (outrx).
func Convert(in []int, inrx int, outrx int) (out []int, err error) {
	if inrx < 2 {
		return nil, errors.New("inrx must be greater than 1")
	}
	if outrx < 2 {
		return nil, errors.New("outrx must be greater than 1")
	}

	if len(in) == 0 {
		return
	}

	outlen := int(math.Ceil(float64(len(in)) * math.Log(float64(inrx)) / math.Log(float64(outrx))))
	out = make([]int, outlen)
	outlen = 0

	for ini, inv := range in {
		if inv >= inrx {
			return nil, fmt.Errorf("in[%v]: %v must be less than inrx: %v", ini, inv, inrx)
		}
		if inv < 0 {
			return nil, fmt.Errorf("in[%v]: %v must be greater or equal 0", ini, inv)
		}

		carry := inv
		outi := 0

		for outi < outlen || carry > 0 {
			outv := out[outi]
			if outv > 0 {
				outv = outv*inrx + carry
			} else {
				outv = carry
			}

			carry = outv / outrx
			out[outi] = outv % outrx

			outi++
		}

		if outi > outlen {
			outlen = outi
		}
	}

	out = out[:outlen]

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return
}

// ConvertBytes converts input bytes (in) with 256 radix
// to output bytes (out) with radix (outrx).
func ConvertBytes(in []byte, outrx int) (out []byte, err error) {
	if outrx < 2 {
		return nil, errors.New("outrx must be greater than 1")
	}
	if outrx > 256 {
		return nil, errors.New("outrx must be less than 256")
	}

	if len(in) == 0 {
		return
	}

	rx := big.NewInt(int64(outrx))
	zero := big.NewInt(0)
	num := new(big.Int).SetBytes(in)

	for num.Cmp(zero) != 0 {
		mod := new(big.Int)
		num.DivMod(num, rx, mod)
		out = append(out, byte(mod.Int64()))
	}

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return
}
