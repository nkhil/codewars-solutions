package main

import (
	"fmt"
	"strings"
)

func main() {
	result := Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"})
	fmt.Printf(fmt.Sprint(result))
}

func Points(games []string) int {
	var totalScore int

	for _, g := range games {
		parts := strings.Split(g, ":")
		myTeamScore := parts[0]
		otherTeamScore := parts[1]

		switch true {
		case myTeamScore > otherTeamScore:
			totalScore += 3
		case otherTeamScore == myTeamScore:
			totalScore += 1
		}
	}

	return totalScore
}

// From Codewars
//
// func Points(games []string) int {
//   result := 0
//   for _, game := range games {
//     str := strings.Split(game, ":")
//     x, y := str[0], str[1]
//     switch {
//       case x > y:
//         result += 3
//       case x == y:
//         result += 1
//     }
//   }
//   return result
// }
