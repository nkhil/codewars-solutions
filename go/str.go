package main

// Trying to understand https://go.dev/blog/strings

import (
	"fmt"
)

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
}
