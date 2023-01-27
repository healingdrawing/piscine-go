package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i, j := 0, len(args)-1; i < j; i, j = i+1, j-1 {
		args[i], args[j] = args[j], args[i]
	}

	for _, arg := range args {
		for _, r := range []rune(arg) {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
