package main

import (
	"fmt"
	"strings"
)

func main() {
	res := Accum("ZpglnRxqenU")
	fmt.Println(res)
}

func Accum(s string) string {
	result := ""
  for i, char := range s {
		for j := 0; j <= i; j++ {
			if j == 0 {
				result += strings.ToUpper(string(char))
			} else {
				result += string(char)
			}
		}

		if i != len(s)-1 {
			result += "-"
	}
	}

	return result
}
