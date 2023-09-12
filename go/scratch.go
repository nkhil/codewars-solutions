package main

import "fmt"

func main() {
	str := "foo"
	byteArray := []byte(str)
	fmt.Println(byteArray)
	backToString := string(byteArray[:])
	fmt.Println(backToString)
}
