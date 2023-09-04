package main

import "fmt"

func main() {
	fmt.Printf(
		fmt.Sprint(Calculate([][2]int{{3,0},{9,1},{4,10},{12,2},{6,1},{7,10}})),
	)
}

func Calculate(busStops [][2]int) int {
	peopleOnBus := 0

	for _, c:= range busStops  {
		peopleOnBus += c[0]
		peopleOnBus -= c[1]
	}

	return peopleOnBus
}
