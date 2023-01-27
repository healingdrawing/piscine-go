package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, r := range word {
			pR(r)
		}
		end()
	}
}

func pR(r rune) {
	z01.PrintRune(r)
}

func end() { z01.PrintRune('\n') }
