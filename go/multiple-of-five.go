package main

import "fmt"

func main () {
	fmt.Printf(fmt.Sprint(
		RoundToNext5(-1),
	))
}
func RoundToNext5(n int) int {
	if (n == 0) {
		return n
	}
  nextMultipleOf5 := (n + 5 - n % 5)
  return nextMultipleOf5
}
