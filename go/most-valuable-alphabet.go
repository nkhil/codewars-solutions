package main

// In this Kata, you will be given a string and your task is to
// return the most valuable character.
// The value of a character is the
// difference between the index of its last occurrence
// and the index of its first occurrence.
// Return the character that has the highest value. If there is a tie, return the alphabetically lowest character. [For Golang return rune]
// All inputs will be lower case.

func Solve(s string) rune {
	// Create a map to store the character and its last occurrence.
	lastOccurrence := make(map[rune]int)

	// Loop through the string.
	for i, c := range s {
		// Update the last occurrence of the character.
		lastOccurrence[c] = i
		// 120: 0
	}

	// Initialize the most valuable character and its value.
	var mostValuable rune
	var mostValuableValue int

	for i, r := range s {
		lastOccurrenceIndex := lastOccurrence[r] // 0
		score := lastOccurrenceIndex - i

		if score == mostValuableValue {
			if mostValuable == 0 {
				mostValuableValue = score
				mostValuable = r
			} else if r < mostValuable {
				mostValuableValue = score
				mostValuable = r
			}
		} else if score > mostValuableValue {
			mostValuableValue = score
			mostValuable = r
		}
	}

	// Return the most valuable character.
	return mostValuable
}

// From Codewars
// func main() {
// 	str := "aabccc"

// 	mostValuableCharacter := Solve(str)

// 	fmt.Println("The most valuable character is", mostValuableCharacter)
// }

// func Solve(s string) rune {
//   indices := map[rune][]int{}
//   maxValue, diff, mostValuable := -1, 0, []rune(s)[0]

//   for i, c := range s {
//     if _, ok := indices[c]; ok {
//       indices[c][1] = i
//     } else {
//       indices[c] = []int{i, i}
//     }
//   }

//   for key, val := range indices {
//     diff = val[1] - val[0]
//     if diff > maxValue {
//       maxValue = diff
//       mostValuable = key
//     } else if diff == maxValue && key < mostValuable {
//       mostValuable = key
//     }
//   }

//   return mostValuable
// }

// import . "strings"

// func Solve(s string) rune {
//   	max,r := 0,rune(123)
//   	for _,c := range s {
//     		d := LastIndex(s, string(c)) - Index(s, string(c))
//     		if max < d || max == d && r > c {
//     			  max = d; r = c
//     		}
//   	}
//   	return r
// }
