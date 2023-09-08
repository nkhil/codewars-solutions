package main

import "fmt"

func main() {
	str := "abx"
	res := []rune("z")

	for _, r := range str {
		fmt.Println(fmt.Sprint(r))

		if r < res[0] {
			res[0] = r
		}
	}
	fmt.Println("---")
	fmt.Println(
		fmt.Sprint(res[0]),
	)
}
