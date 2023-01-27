package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)
	for _, r := range runes {
		if byte(r) < 32 {
			return false
		}
	}
	return true
}
