package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	lena := len(args)
	if lena == 1 {
		rc := os.Stdin
		outar, err := io.ReadAll(rc)
		if err == nil {
			for _, r := range outar {
				pr(rune(r))
			}
		}
	} else {
		files := args[1:]
		firstStep := true // to simulate weird task output
		for _, fname := range files {
			bytes, err := ioutil.ReadFile(fname)
			if err == nil {
				stringPrinter(string(bytes))
			} else {
				if firstStep {
					stringPrinter("ERROR: " + err.Error() + "\n")
					os.Exit(1)
					firstStep = false
				} else {
					stringPrinter(err.Error() + "\n")
				}
			}
		}
	}
}

func printEmptyCall() {
	s := ""
	// s := "Hello\nHello\n^C\n"
	stringPrinter(s)
}

func stringPrinter(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func pr(r rune) { z01.PrintRune(r) }
