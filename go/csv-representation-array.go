/*
https://www.codewars.com/kata/5a34af40e1ce0eb1f5000036/train/go
Create a function that returns the CSV representation of a two-dimensional numeric array.

Example:

input:
   [[ 0, 1, 2, 3, 4 ],
    [ 10,11,12,13,14 ],
    [ 20,21,22,23,24 ],
    [ 30,31,32,33,34 ]]

output:
     '0,1,2,3,4\n'
    +'10,11,12,13,14\n'
    +'20,21,22,23,24\n'
    +'30,31,32,33,34'
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	arrayOfArrays := [][]int{numbers}

	fmt.Println(
		ToCsvText(arrayOfArrays),
	)
}
// [[1, 2, 3], [4, 5, 6]]
func ToCsvText(data [][]int) string {
	var b strings.Builder
    for _, row := range data {
			for i, cell := range row {
					b.WriteString(strconv.Itoa(cell))
					if i < len(row) - 1 {
						b.WriteString(",")
					}
			}
			b.WriteString("\n")
    }
    return b.String()
}
