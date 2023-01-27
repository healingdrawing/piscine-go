package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	intNum := atoiBase(nbr, baseFrom)
	return printNbrBase(intNum, baseTo)
}

func atoiBase(s string, base string) int {
	if isBaseOk(base) {
		var indices []int
		for _, r := range []rune(s) {
			for ind, br := range []rune(base) {
				if r == br {
					indices = append(indices, ind)
				}
			}
		}
		sum := 0
		for i := 0; i < len(indices); i++ {
			sum *= len(base)
			sum += indices[i]
		}
		return sum
	}
	return 0
}

/* base unique and greater than 1 rune. (at least two runes)*/
func isBaseOk(base string) bool {
	var uniqueRunes []rune
	baseRunes := []rune(base)
	if len(baseRunes) < 2 {
		return false
	}
	for _, r := range baseRunes {
		for _, ur := range uniqueRunes {
			if r == ur {
				return false
			}
		}
		uniqueRunes = append(uniqueRunes, r)
	}
	// check base for + - which is not allowed
	for _, r := range uniqueRunes {
		if r == '+' || r == '-' {
			return false
		}
	}
	return true
}

func printNbrBase(nbr int, base string) string {
	// final number runes
	var frunes []rune

	if isBaseOk(base) {
		minus := nbr < 0
		baseRunes := []rune(base)
		b := len(baseRunes)
		var indices []int
		for nbr != 0 {
			tail := nbr % b
			indices = append(indices, tail)
			nbr = nbr / b
		}
		for i, j := 0, len(indices)-1; i < j; i, j = i+1, j-1 {
			indices[i], indices[j] = indices[j], indices[i]
		}

		if minus {
			frunes = append(frunes, '-')
			// RP('-')
		}
		for _, v := range indices {
			frunes = append(frunes, baseRunes[posIndex(v)])
			// RP(baseRunes[posIndex(v)])
		}
	} else {
		// RP('N')
		// RP('V')
	}
	return string(frunes)
	// RP('\n')
}

// positive index
func posIndex(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
