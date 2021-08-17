package radix_test

import (
	"fmt"
	"log"

	"github.com/bergusman/radix-go"
)

func Example() {
	out, err := radix.Convert([]int{1, 3, 3, 7}, 10, 16)
	if err != nil {
		log.Fatal(err)
	}

	str, err := radix.Encode(out, "0123456789ABCDEF")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str)
	// Output: 539
}

func Example_emoji() {
	out, err := radix.Convert([]int{1, 3, 3, 7}, 10, 5)
	if err != nil {
		log.Fatal(err)
	}

	str, err := radix.Encode(out, "🌑🌘🌗🌖🌕")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str)
	// Output: 🌗🌑🌖🌗🌗
}

func Example_bitcoin() {
	addr := []byte{138, 238, 64, 184, 232, 126, 176, 91, 195, 185, 255, 144, 35, 73, 187, 45, 209, 154, 94, 144}
	str, err := radix.Base58Encode(addr, radix.AlphabetBitcoin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str)
	// Output: 2wG9ewuHafzR4yG2hYu8U3YmpZxb
}

func ExampleConvert() {
	out, err := radix.Convert([]int{1, 3, 3, 7}, 10, 2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
	// Output: [1 0 1 0 0 1 1 1 0 0 1]
}

func ExampleEncode() {
	out, err := radix.Convert([]int{1, 3, 3, 7}, 10, 2)
	if err != nil {
		log.Fatal(err)
	}

	str, err := radix.Encode(out, "01")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str)
	// Output: 10100111001
}
