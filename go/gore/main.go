package main

import (
	"os"

	// "strings"
	// "strconv"
	"fmt"
)

func main() {
	args := os.Args
	fmt.Println("hello")
	fmt.Println(args)

	lena := len(args)
	if lena == 3 {
		good(args)
	}
}

// cut broken rune at the end of string, was detected before, probably no needed at the moment
func cutLastRune(str string) (cutted string) {
	runes := []rune(str)
	runes = runes[:len(runes)-1]
	cutted = string(runes)
	return
}

// read file from "fpath" location and return (string, bool)
func rfile(fpath string) (str string, ok bool) {
	bytes, err := os.ReadFile(fpath)
	if err == nil {
		str = cutLastRune(string(bytes))
		ok = true
		fmt.Println(string(str))
	} else {
		str = ""
		ok = false
	}
	return
}

// input data looks good, continue calculation
func good(args []string) {
	fmt.Println(args)

}
