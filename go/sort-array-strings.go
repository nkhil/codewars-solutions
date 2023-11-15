// Write a function that takes an array of strings as an argument and returns a sorted array containing the same strings, ordered from shortest to longest.

// For example, if this array were passed as an argument:

// ["Telescopes", "Glasses", "Eyes", "Monocles"]

// Your function would return the following array:

// ["Eyes", "Glasses", "Monocles", "Telescopes"]

// All of the strings in the array passed to your function will be different lengths, so you will not have to decide how to order multiple strings of the same length.
package main

import "fmt"

func main() {
	fmt.Println(SortByLength([]string{"Telescopes", "Glasses", "Eyes", "Monocles"}))
}

func SortByLength(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if len(arr[j]) > len(arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

// From codewars:
// func SortByLength(arr []string) []string {
//   newArr := arr[:]
//   sort.Slice(newArr, func(i, j int) bool { return len(newArr[i]) < len(newArr[j]) })
//   return newArr
// }
