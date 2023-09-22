// There is an array with some numbers. All numbers are equal except for one. Try to find it!
// findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
// findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55
// It’s guaranteed that array contains at least 3 numbers.

package main

import "fmt"

func FindUniq(arr []float32) float32 {
  var result float32

  // Handle when the unique value is the first element
  if arr[0] != arr[1] && arr[1] == arr[2] { return arr[0] }

  for i, v := range arr {
    if i == 0 {
      result = v
      } else if result != v {
      result = v
      break
    }
  }

  return result
}

func main() {
  fmt.Println(FindUniq([]float32{ 1, 1, 1, 2, 1, 1 }))
  fmt.Println(FindUniq([]float32{ 2, 1, 1, 1, 1, 1 }))
  fmt.Println(FindUniq([]float32{ 0, 0, 0.55, 0, 0 }))
}
