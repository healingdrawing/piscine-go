package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		upperFlag := args[1] == "--upper"

		tup := []rune{' ', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 56: ' '}
		tlo := []rune{' ', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 56: ' '}

		args = MakeArgsCorrect(args, upperFlag)

		// startindex := 1
		// if upperFlag {
		// 	startindex = 2
		// }
		for i := 0; i < len(args); i++ {
			letterindex := int(Stoi(args[i]))
			if upperFlag {
				z01.PrintRune(tup[letterindex])
			} else {
				z01.PrintRune(tlo[letterindex])
			}
		}
		z01.PrintRune('\n')
	}
}

// filtering + replacing all failed incoming by space as in example
func MakeArgsCorrect(args []string, upperFlag bool) []string {
	startindex := 1
	if upperFlag {
		startindex = 2
	}
	var filteredArgsTail []string
	for i := startindex; i < len(args); i++ {
		arg := args[i]
		argi := Stoi(arg)
		if argi > -1 && argi < 27 {
			filteredArgsTail = append(filteredArgsTail, arg)
		} else {
			filteredArgsTail = append(filteredArgsTail, "0")
		}

	}
	return filteredArgsTail
}

// check that arguments only numbers
func IsArgsCorrect(args []string, upperFlag bool) bool {
	startindex := 1
	if upperFlag {
		startindex = 2
	}
	for i := startindex; i < len(args); i++ {
		if !OnlyNumbersInside(args[i]) {
			return false
		}
	}
	return true
}

// check one argument is only number
func OnlyNumbersInside(arg string) bool {
	for _, r := range arg {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// string number to integer
func Stoi(s string) int {
	sar := []rune(s)
	ln := len(sar)
	numsum := 0
	for i := 0; i < ln; i++ {
		numsum = numsum*10 + (int(sar[i] - '0'))
	}
	return numsum
}
