package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	// check minus
	negative := false
	if n < 0 {
		negative = true
	}
	if negative {
		z01.PrintRune('-')
	}
	ra, imax := ITI(n)
	for i := 0; i <= imax; i++ {
		if n < 0 {
			z01.PrintRune(IRCLONE(-ra[i]))
		} else {
			z01.PrintRune(IRCLONE(ra[i]))
		}
	}
}

func Reverse(s [100]int, imax int) [100]int {
	var r [100]int
	for i := imax; i > -1; i-- {
		r[imax-i] = s[i]
	}
	return r
}

func ITI(n int) (s [100]int, index int) {
	i := 0
	for n > 10 {
		t := n % 10
		n = n / 10
		s[i] = t
		i++
	}

	for n < -10 {
		t := n % 10
		n = n / 10
		s[i] = t
		i++
	}

	s[i] = n
	index = i

	s = Reverse(s, index)
	return
}

func IRCLONE(s int) (r rune) {
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
