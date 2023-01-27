package piscine

func AtoiBase(s string, base string) int {
	if IsBaseOk(base) {
		var indices []int
		// iterate runes of incoming number representatoin of string
		for _, r := range []rune(s) {
			// iterate runes of used base for number
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
func IsBaseOk(base string) bool {
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
