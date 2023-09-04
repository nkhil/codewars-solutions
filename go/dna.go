package main

import "fmt"

func main () {
	fmt.Printf(getComplementaryDNA("GTAT"))
}

func getComplementaryDNA(dna string) string {
  complements := map[rune]rune{
    'A': 'T',
    'T': 'A',
    'C': 'G',
    'G': 'C',
  }

  var rcDNA []rune
  for _, c := range dna {
    rcDNA = append(rcDNA, complements[c])
  }

  return string(rcDNA)
}
