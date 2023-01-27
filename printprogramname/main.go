package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	runes := []rune(os.Args[0])
	slashIndex := -1
	for i, r := range runes {
		if r == os.PathSeparator {
			slashIndex = i
		}
	}
	for i := slashIndex + 1; i < len(runes); i++ {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}
