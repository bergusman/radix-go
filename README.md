# Radix (Base) Conversion

Sequence of integers (`byte`, `int`) can be represent big number in positional numeral system with specified radix.

This project's functions convert input sequence from one radix to output sequence with other radix. Also encode/decode these sequences with specified alphabet.

> Aim of this project to implement universal radix converter. But implementation of `[]byte` input conversion using `big.Int` is more efficient than universtal implementation of `[]int` input conversion. Check benchmarks into `convert_test.go`.

Base58 encoding (used by Bitcoin and others) builds on this conversion of big number from base 256 to base 58.

### Examples

```Go
package main

import (
	"fmt"
	"log"
	
	"github.com/bergusman/radix-go"
)

func main() {
	out, err := radix.Convert([]int{1, 3, 3, 7}, 10, 16)
	if err != nil {
		log.Fatal(err)
	}

	str, err := radix.Encode(out, "0123456789ABCDEF")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str) // Output: 539

	out, err = radix.Convert([]int{1, 3, 3, 7}, 10, 5)
	if err != nil {
		log.Fatal(err)
	}

	str, err = radix.Encode(out, "🌑🌘🌗🌖🌕")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str) // Output: 🌗🌑🌖🌗🌗
}
```

#### Base58

```Go
package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/bergusman/radix-go"
)

func main() {
	addr, err := hex.DecodeString("8aee40b8e87eb05bc3b9ff902349bb2dd19a5e90")
	if err != nil {
		log.Fatal(err)
	}

	str, err := radix.Base58Encode(addr, radix.AlphabetBitcoin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str) // Output: 2wG9ewuHafzR4yG2hYu8U3YmpZxb
}
```

Can check here: http://lenschulwitz.com/base58
