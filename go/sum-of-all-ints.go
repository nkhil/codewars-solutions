//Your task in this kata is to implement a function that calculates the sum of the integers inside a string.
// For example, in the string "The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog", the sum of the integers is 3635.

package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Printf(
		fmt.Sprint(SumOfIntegersInString("10a10a3zz")),
	)
}
func SumOfIntegersInString(strng string) int {
  r := 0
  for _, s := range regexp.MustCompile(`\d+`).FindAllString(strng, -1) {
    x, _ := strconv.Atoi(s)
    r += x
  }
  return r
}