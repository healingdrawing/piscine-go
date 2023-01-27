package piscine

import (
	"github.com/01-edu/z01"
)

var (
	hasQueen [8][8]bool
	inAttack [8][8]int
)

func placeQueen(i, j int) {
	hasQueen[i][j] = true
	// update attack counts
	// row
	for c := 0; c < 8; c++ {
		inAttack[i][c]++
	}
	// col
	for r := 0; r < 8; r++ {
		inAttack[r][j]++
	}
	inAttack[i][j] -= 2 // the (i,j) cell has been counted twice in the above iterations
	// diagonals
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if r == i && c == j {
				inAttack[r][c]++
				continue
			}
			if r-i == c-j || r-i == -(c-j) {
				inAttack[r][c]++
			}
		}
	}
}

func removeQueen(i, j int) {
	hasQueen[i][j] = false
	// update attack counts
	// row
	for c := 0; c < 8; c++ {
		inAttack[i][c]--
	}
	// col
	for r := 0; r < 8; r++ {
		inAttack[r][j]--
	}
	inAttack[i][j] += 2
	// diagonals
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if r == i && c == j {
				inAttack[r][c]--
				continue
			}
			if r-i == c-j || r-i == -(c-j) {
				inAttack[r][c]--
			}
		}
	}
}

func solve(r int) {
	if r == 7 {
		// can we place the queen somewhere
		for c, cnt := range inAttack[7] {
			if cnt == 0 {
				placeQueen(r, c)
				LinePrinter(hasQueen)
				removeQueen(r, c)
			}
		}
	} else {
		for c, cnt := range inAttack[r] {
			if cnt == 0 {
				placeQueen(r, c)
				solve(r + 1)      // solve the next row (recursive)
				removeQueen(r, c) // remove this queen (backtrack)
			}
		}
	}
}

var rarr = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func LinePrinter(data [8][8]bool) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if data[i][j] {
				z01.PrintRune(rarr[j+1])
				break
			}
		}
	}
	z01.PrintRune('\n')
}

func EightQueens() {
	solve(0)
}
