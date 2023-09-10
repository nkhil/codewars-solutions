// A triangle is called an equable triangle if its area equals its perimeter.
// Return true, if it is an equable triangle, else return false.
// You will be provided with the length of sides of the triangle.

package main

import (
	"fmt"
	"math"
)

func EquableTriangle(a, b, c int) bool {
	perimeter := a + b + c
	semiPerimeter := perimeter / 2

	areaFloat := math.Sqrt(float64(semiPerimeter * (semiPerimeter - a) * (semiPerimeter - b) * (semiPerimeter - c)))
	area := int(math.Floor(areaFloat))

	return area == perimeter
}

func main() {
	fmt.Printf(fmt.Sprint(EquableTriangle(5, 12, 13)))
	fmt.Printf(fmt.Sprint(EquableTriangle(2, 3, 4)))
}
