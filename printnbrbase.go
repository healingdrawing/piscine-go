package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if IsBaseGood(base) {
		minus := nbr < 0
		baseRunes := []rune(base)
		b := len(baseRunes)
		var indices []int
		for nbr != 0 {
			tail := nbr % b
			indices = append(indices, tail)
			nbr = nbr / b
		}
		for i, j := 0, len(indices)-1; i < j; i, j = i+1, j-1 {
			indices[i], indices[j] = indices[j], indices[i]
		}
		if minus {
			RP('-')
		}
		for _, v := range indices {
			RP(baseRunes[PosIndex(v)])
		}
	} else {
		RP('N')
		RP('V')
	}
	// RP('\n')
}

// positive index
func PosIndex(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

/* base unique and greater than 1 rune. (at least two runes)*/
func IsBaseGood(base string) bool {
	var uniqueRunes []rune
	baseRunes := []rune(base)
	if len(baseRunes) < 2 {
		return false
	}
	for _, r := range baseRunes {
		for _, ur := range uniqueRunes {
			if r == ur {
				return false
			}
		}
		uniqueRunes = append(uniqueRunes, r)
	}
	// check base for + - which is not allowed
	for _, r := range uniqueRunes {
		if r == '+' || r == '-' {
			return false
		}
	}
	return true
}

func RP(r rune) { z01.PrintRune(r) }
