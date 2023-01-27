package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	rlen := len(runes)
	if n > 0 && n < rlen+1 {
		return runes[n-1]
	} else {
		return 0
	}
}
