package piscine

func IsLower(s string) bool {
	t := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
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
