// https://www.codewars.com/kata/558fa34727c2d274c10000ae/solutions/go
// Write a program that, given a word, computes the scrabble score for that word.

// Letter Values
// You'll need these:

// Letter                           Value
// A, E, I, O, U, L, N, R, S, T       1
// D, G                               2
// B, C, M, P                         3
// F, H, V, W, Y                      4
// K                                  5
// J, X                               8
// Q, Z                               10
// There will be a preloaded map DictScores with all these values: DictScores["E"] == 1

// Examples
// "cabbage" --> 14
// "cabbage" should be scored as worth 14 points:

// 3 points for C
// 1 point for A, twice
// 3 points for B, twice
// 2 points for G
// 1 point for E
// And to total:

// 3 + 2*1 + 2*3 + 2 + 1 = 3 + 2 + 6 + 3 = 14

// Empty string should return 0. The string can contain spaces and letters (upper and lower case), you should calculate the scrabble score only of the letters in that string.

// ""           --> 0
// "STREET"    --> 6
// "st re et"    --> 6
// "ca bba g  e" --> 14

const scores = {
	' ': 0, // blank
	'A': 1, 'a': 1,
	'E': 1, 'e': 1,
	'I': 1, 'i': 1,
	'O': 1, 'o': 1,
	'U': 1, 'u': 1,
	'L': 1, 'l': 1,
	'N': 1, 'n': 1,
	'R': 1, 'r': 1,
	'S': 1, 's': 1,
	'T': 1, 't': 1,
	'D': 2, 'd': 2,
	'G': 2, 'g': 2,
	'B': 3, 'b': 3,
	'C': 3, 'c': 3,
	'M': 3, 'm': 3,
	'P': 3, 'p': 3,
	'F': 4, 'f': 4,
	'H': 4, 'h': 4,
	'V': 4, 'v': 4,
	'W': 4, 'w': 4,
	'Y': 4, 'y': 4,
	'K': 5, 'k': 5,
	'J': 8, 'j': 8,
	'X': 8, 'x': 8,
	'Q': 10, 'q': 10,
	'Z': 10, 'z': 10,
}

function score(word) {
  const wordArr = word.split('')
  return wordArr.reduce((scr, letter) => {
    scr += scores[letter]
    return scr
  }, 0)
}
