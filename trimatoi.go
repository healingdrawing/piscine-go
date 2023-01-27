package piscine

func TrimAtoi(s string) int {
	te := []rune{'-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var fr []rune
	for _, c := range s {
		for _, tec := range te {
			if c == tec {
				fr = append(fr, c)
			}
		}
	}
	var finalRunes []rune
	haveMinus := false
	// check the minus is first symbol before numbers
	if len(fr) > 1 && fr[0] == '-' {
		finalRunes = append(finalRunes, '-')
		fr = fr[1:]
		haveMinus = true
	} else if len(fr) == 1 && fr[0] == '-' {
		fr = []rune{'0'}
	}
	// add all other
	for _, r := range fr {
		if r != '-' {
			finalRunes = append(finalRunes, r)
		}
	}
	// join runes into string
	finalString := ""
	for _, r := range finalRunes {
		finalString += string(r)
	}

	var startindex int
	if haveMinus {
		startindex = 1
	} else {
		startindex = 0
	}
	numsum := 0
	for i := startindex; i < len(finalRunes); i++ {
		numsum = numsum*10 + (int(finalRunes[i] - '0'))
	}
	// reproduce negative sign
	if haveMinus {
		numsum *= -1
	}
	return numsum
}
