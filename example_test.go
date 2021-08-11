package radix

import (
	"fmt"
	"log"
)

func Example() {
	out, err := Convert([]int{1, 3, 3, 7}, 10, 16)
	if err != nil {
		log.Fatal(err)
	}

	str, err := Encode(out, "0123456789ABCDEF")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str)
	// Output: 539
}

func Example_emoji() {
	out, err := Convert([]int{1, 3, 3, 7}, 10, 5)
	if err != nil {
		log.Fatal(err)
	}

	str, err := Encode(out, "🌑🌘🌗🌖🌕")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str)
	// Output: 🌗🌑🌖🌗🌗
}
