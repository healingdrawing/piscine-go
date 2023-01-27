package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, arg := range os.Args[1:] {
		for _, r := range []rune(arg) {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
