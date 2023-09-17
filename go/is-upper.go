package main

import (
	"fmt"
	"strings"
	"unicode"
)

type MyString string

func CheckWord(s MyString) bool {
	runes := []rune(s)
	result := false

	for _, r := range runes {
		if unicode.IsUpper(r) {
			result = true
		} else {
			result = false
			break
		}
	}

	return result
}

func (s MyString) IsUpperCase() bool {
	isUpper := false
	removedSpaces := strings.ReplaceAll(string(s), " ", "")

	fmt.Println(removedSpaces)

	for _, w := range removedSpaces {
		if CheckWord(MyString(w)) {
			isUpper = true
		} else {
			isUpper = false
			break
		}
	}

	return isUpper
}

func main() {
	fmt.Println(
		MyString("O ").IsUpperCase(),
	)
}

// From CodeWars
// func (s MyString) IsUpperCase() bool {
// Your code here!
// return s == MyString(strings.ToUpper(string(s)))
// }
