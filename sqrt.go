package piscine

func Sqrt(nb int) int {
	t := TemplateCreator(DigitsNumber(nb))
	solved := false
	solution := 0
	for !solved {
		tt := t * t
		if tt <= nb {
			if tt == nb {
				solved = true
				solution = t
			} else {
				solved = true
				solution = 0
			}
		} else {
			t--
		}
	}
	return solution
	// return solution, nb, DigitsNumber(nb), TemplateCreator(DigitsNumber(nb))
}

func TemplateCreator(dn int) int {
	t := 1
	for i := 0; i < dn; i++ {
		t *= 10
	}
	return t
}

func DigitsNumber(nb int) int {
	dn := 0
	for nb > 0 {
		nb /= 10
		dn++
	}
	return HalfDigitsNumber(dn)
}

func HalfDigitsNumber(dn int) int {
	extra := 0
	if dn%2 > 0 {
		extra = 1
	}
	return dn/2 + extra
}
