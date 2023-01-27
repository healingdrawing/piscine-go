package piscine

func ToUpper(s string) string {
	lower := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	upper := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	runes := []rune(s)
	for rindex, r := range runes {
		isLower, index := replacerIndex(r, lower)
		if isLower {
			runes[rindex] = upper[index]
		}
	}
	return string(runes)
}

func replacerIndex(r rune, template []rune) (bool, int) {
	for i, tr := range template {
		if r == tr {
			return true, i
		}
	}
	return false, -1
}
