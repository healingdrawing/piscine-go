package piscine

func Swap(a *int, b *int) {
	c := *a + 0
	*a = *b + 0
	*b = c
}
