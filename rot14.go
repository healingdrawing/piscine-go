package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for ind, r := range runes {
		if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' {
			runes[ind] = arot(r, 14)
		}
	}
	return string(runes)
}

func arot(r rune, shift int) rune {
	teBig := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	te := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	var big bool
	var rotind int
	for ind, ter := range teBig {
		if ter == r {
			rotind = ind + shift
			if rotind > len(teBig)-1 {
				rotind = rotind - len(teBig)
			}
			big = true
			break
		}
	}
	if big {
		return teBig[rotind]
	}
	for ind, ter := range te {
		if ter == r {
			rotind = ind + shift
			if rotind > len(te)-1 {
				rotind = rotind - len(te)
			}
			break
		}
	}
	return te[rotind]
}
