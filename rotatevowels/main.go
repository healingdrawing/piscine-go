package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		nLine()
	} else {
		args = args[1:]
		sum := "" // all in one string
		for ind, arg := range args {
			if ind > 0 {
				sum += " "
			}
			sum += arg
		}
		sum += " " // looks like fail from system side
		// arr of runes
		runes := []rune(sum)
		var vowelsInd []int
		var vowelsRune []rune
		// collect indices and runes
		for ind, r := range runes {
			if isVowel(r) {
				vowelsInd = append(vowelsInd, ind)
				vowelsRune = append(vowelsRune, r)
			}
		}
		// switching
		for a, z := 0, len(vowelsInd)-1; a < z; a, z = a+1, z-1 {
			leftIndex := vowelsInd[a]
			rightIndex := vowelsInd[z]
			runes[leftIndex], runes[rightIndex] = runes[rightIndex], runes[leftIndex]
		}
		pRunes(runes)
	}
}

func isVowel(r rune) bool {
	te := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, ter := range te {
		if ter == r {
			return true
		}
	}
	return false
}

// print arr/slice of runes
func pRunes(ar []rune) {
	for _, r := range ar {
		pR(r)
	}
	nLine()
}

// print rune
func pR(r rune) {
	z01.PrintRune(r)
}

// print new line rune
func nLine() { pR('\n') }
