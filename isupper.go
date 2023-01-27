package piscine

func IsUpper(s string) bool {
	t := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	runes := []rune(s)
	for _, r := range runes {
		result := false

		for _, tr := range t {
			if r == tr {
				result = true
			}
		}
		if result == false {
			return false
		}
	}
	return true
}
