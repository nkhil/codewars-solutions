package main

import (
	"fmt"
)

func main() {
	// fmt.Println(reverseWordsInString("double  spaced  words"))
	reverseWord("hello")
	// fmt.Println(reverseWord("world"))
}

// func reverseWordsInString(str string) string {
// 	var result string = ""
// 	words := strings.Split(str, " ")
// 	for _, word := range words {
// 		result += reverseWord(word) + " "
// 	}
// 	return result
// }

func reverseWord(str string) string {
	var result string = ""
  for _, v := range str {
		fmt.Printf("%c",v)
		result = string(v) + result
	}

	return result
}
