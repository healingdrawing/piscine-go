package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	printail := true
	for a := 0; a < 100; a++ {
		for b := a; b < 100; b++ {
			if a < b {
				if a == 0 && b == 1 {
					printail = false
				}
				if printail {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				if a < 10 {
					z01.PrintRune(rune(IR(0)))
					z01.PrintRune(rune(IR(a)))
				} else {
					tail := a % 10
					z01.PrintRune(rune(IR((a - tail) / 10)))
					z01.PrintRune(rune(IR(tail)))
				}

				z01.PrintRune(' ')
				if b < 10 {
					z01.PrintRune(rune(IR(0)))
					z01.PrintRune(rune(IR(b)))
				} else {
					tail := b % 10
					z01.PrintRune(rune(IR((b - tail) / 10)))
					z01.PrintRune(rune(IR(tail)))
				}

				printail = true

			}
		}
	}
	z01.PrintRune('\n')
}

func IR2(s int) (r rune) {
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
