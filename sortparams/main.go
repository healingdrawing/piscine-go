package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	args = SortStrArrAscending(args)

	for _, arg := range args {
		for _, r := range []rune(arg) {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

func SortStrArrAscending(aof []string) []string {
	lena := len(aof) - 1
	sortedPairs := 0
	for sortedPairs < lena {
		for i := 0; i < lena; i++ {
			if aof[i] > aof[i+1] {
				sortedPairs = 0
				aof[i], aof[i+1] = aof[i+1], aof[i]
			} else {
				sortedPairs++
			}
		}
	}
	return aof
}
