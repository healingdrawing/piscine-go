package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// integer to runes
func itor(n int) []rune {
	te := map[int]rune{0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9'}
	var arune []rune
	for n > 0 {
		arune = append(arune, rune(te[n%10]))
		n /= 10
	}
	for i, j := 0, len(arune)-1; i < j; i, j = i+1, j-1 {
		arune[i], arune[j] = arune[j], arune[i]
	}
	return arune
}

func printer(p *point) {
	sx := "x = "
	sy := ", y = "
	var runes []rune
	runes = append(runes, []rune(sx)...)
	runes = append(runes, itor(p.x)...)
	runes = append(runes, []rune(sy)...)
	runes = append(runes, itor(p.y)...)
	for _, r := range runes {
		pr(r)
	}
	pr('\n')
}

func pr(r rune) {
	z01.PrintRune(r)
}

func main() {
	points := &point{}
	setPoint(points)
	printer(points)
}
