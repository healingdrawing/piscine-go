package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	lena := len(args)
	fname := "quest8.txt"
	if lena < 2 {
		fmt.Println("File name missing")
	} else if lena > 2 {
		fmt.Println("Too many arguments")
	} else {
		bytes, err := ioutil.ReadFile(fname)
		if err == nil {
			str := string(bytes)
			runes := []rune(str)
			runes = runes[:len(runes)-1]
			str = string(runes)
			fmt.Println(string(str))
		}
	}
}
