package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := StringToNumber("-123")
	fmt.Println(res)
}

func StringToNumber(str string) int {
	r, err := strconv.Atoi(str)

	if err == nil {
		return r
	}

	return 0
}
