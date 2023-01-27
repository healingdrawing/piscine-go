package piscine

func AlphaCount(s string) int {
	runes := []rune(s)
	counter := 0
	for _, r := range runes {
		if IsLatLetter(r) {
			counter++
		}
	}
	return counter
}

var template = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

func IsLatLetter(r rune) bool {
	for _, temprune := range template {
		if r == temprune {
			return true
		}
	}
	return false
}
