// Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.

// Return your answer as a number.

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func SumMix(array []any) int {
	sum := 0
	for _, value := range array {
		fmt.Println("value is", value)
		fmt.Println("sum", sum)

		switch reflect.TypeOf(value).Kind() {
		case reflect.String:
			num, err := strconv.Atoi(value.(string))
			fmt.Println(num, err)
			if err == nil {
				sum += num
			}
		case reflect.Int:
			sum += value.(int)
		}
	}
	return sum
}

func main() {
	array := []any{8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7}
	sum := SumMix(array)
	fmt.Println("The sum of the array is", sum)
}

// import "strconv"

// func SumMix(arr []any) int {

//   sum := 0

//   for _, val := range arr {

//    switch val := val.(type){
//      case int:
//       sum += val
//      case string:
//       k,_ := strconv.Atoi(val)
//       sum += k
//    }

//   }

//   return sum
// }
