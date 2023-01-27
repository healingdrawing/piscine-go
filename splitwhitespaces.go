package piscine

func SplitWhiteSpaces(s string) []string {
	lena := len(s) - 1
	word := ""
	var words []string
	for ind, r := range s {
		if r == ' ' {
			if word != "" {
				words = append(words, word)
			}
			word = ""
		} else if ind == lena {
			if r != ' ' {
				word += string(r)
			}
			if word != "" {
				words = append(words, word)
			}
		} else {
			word += string(r)
		}
	}
	return words
}
