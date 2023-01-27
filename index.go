package piscine

func Index(s string, toFind string) int {
	index := -1
	runes := []rune(s)
	frunes := []rune(toFind)
	runesl := len(runes)
	frunesl := len(frunes)
	steps := 1 + runesl - frunesl
	if steps < 1 {
		return index
	}
	// start position for comparing , symbol sequences
	for startposition := 0; startposition < steps; startposition++ {
		index = startposition
		for rindex := 0; rindex < frunesl; rindex++ {
			if runes[rindex+startposition] != frunes[rindex] {
				index = -1
				break
			}
		}
		if index > -1 {
			break
		}
	}
	return index
}
