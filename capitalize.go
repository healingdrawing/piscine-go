package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	wasSpace := true
	wasFirstLetter := false
	for ind, r := range runes {
		if isSpaceRune(r) { /*space*/
			wasFirstLetter = false
			wasSpace = true
		} else if isAlphaNumericRune(r) {
			if wasSpace { /*need capitalize*/
				runes[ind] = ToUpperRune(r)
				wasFirstLetter = true
				wasSpace = false
			} else if wasFirstLetter { /*need lowercase*/
				runes[ind] = ToLowerRune(r)
			}
		}
	}
	return string(runes)
}

// space detection section
func isSpaceRune(r rune) bool {
	return !RuneAZ(r) && !Runeaz(r) && !RuneDigit(r)
}

func isAlphaNumericRune(r rune) bool {
	return !isSpaceRune(r)
}

func RuneAZ(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func Runeaz(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func RuneDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func ToUpperRune(r rune) rune {
	lower := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	upper := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	for ind, lowerune := range lower {
		if lowerune == r {
			return upper[ind]
		}
	}
	return r
}

func ToLowerRune(r rune) rune {
	lower := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	upper := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	for ind, uperune := range upper {
		if uperune == r {
			return lower[ind]
		}
	}
	return r
}
