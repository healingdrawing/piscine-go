package main

import (
	"fmt"
	"os"
)

var errors = false

func main() {
	if len(os.Args) == 4 {
		readFile(os.Args[3], true, true)
	} else if len(os.Args) > 4 {
		for i, file := range os.Args[3:] {
			if i == 0 {
				readFile(file, false, true)
			} else {
				readFile(file, false, false)
			}
		}
	}
	if errors {
		os.Exit(1)
	}
}

func atoi(s string) int {
	su := 0
	for _, r := range s {
		su = su*10 + int(r-'0')
	}
	return su
}

func readFile(filename string, solo, first bool) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err.Error() + "\n")
		errors = true
		return
	}
	if !first {
		fmt.Print("\n")
	}
	if !solo {
		fmt.Print("==> " + string(filename) + " <==\n")
	}
	start := len(data) - atoi(os.Args[2])
	if start < 0 {
		start = 0
	}
	fmt.Print(string(data[start:]))
}
