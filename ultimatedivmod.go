package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a + 0
	d := *b + 0
	*a = c / d
	*b = c % d
}
