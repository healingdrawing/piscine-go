package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		switch arg {
		case "01", "galaxy", "galaxy 01":
			fmt.Println("Alert!!!")
			return
		}
	}
}
