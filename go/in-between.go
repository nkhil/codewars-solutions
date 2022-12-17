/*

https://www.codewars.com/kata/55ecd718f46fba02e5000029/train/go

Complete the function that takes two integers (a, b, where a < b) and return an array of all integers between the input parameters, including them.

For example:

a = 1
b = 4
--> [1, 2, 3, 4]
*/

package main

func main() {
	between(2, 20)
}

func between(a, b int) []int {
	result := []int{}
	for i := a; i <= b; i++ {
		result = append(result, i)
	}
	return result
}
