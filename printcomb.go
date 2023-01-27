package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	printail := true
	for a := 0; a < 10; a++ {
		for b := a; b < 10; b++ {
			for c := b; c < 10; c++ {
				if a < b && b < c {
					if a == 0 && b == 1 && c == 2 {
						printail = false
					}
					if printail {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(rune(IR(a)))
					z01.PrintRune(rune(IR(b)))
					z01.PrintRune(rune(IR(c)))

					printail = true

				}
			}
		}
	}
	z01.PrintRune('\n')
}

func IR(s int) (r rune) {
	if s == 0 {
		r = '0'
	} else if s == 1 {
		r = '1'
	} else if s == 2 {
		r = '2'
	} else if s == 3 {
		r = '3'
	} else if s == 4 {
		r = '4'
	} else if s == 5 {
		r = '5'
	} else if s == 6 {
		r = '6'
	} else if s == 7 {
		r = '7'
	} else if s == 8 {
		r = '8'
	} else if s == 9 {
		r = '9'
	}
	return
}
