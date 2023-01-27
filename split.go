package piscine

func Split(s, sep string) []string {
	seprunes := []rune(sep)
	seplen := len(seprunes)
	srunes := []rune(s)
	slen := len(srunes)
	var sepsIndices []int
	// add first index as natural separator minus length of separator
	sepsIndices = append(sepsIndices, 0-seplen)
	for ind := range s {
		if ind < 1+slen-seplen { // there is still enough positions for separator in string
			sepFound := false
			for sind, rind := 0, ind; sind < seplen; sind, rind = sind+1, rind+1 {
				if seprunes[sind] == srunes[rind] {
					sepFound = true
				} else {
					sepFound = false
					break
				}
			}
			if sepFound { // separator found, add his index
				sepsIndices = append(sepsIndices, ind)
			}
		}
	}
	var result []string
	// slicing srunes into parts from separator to next separator or end of string
	var leftBorder int
	var rightBorder int
	for i := 0; i < len(sepsIndices); i++ {
		if i < len(sepsIndices)-1 {
			leftBorder = sepsIndices[i] + seplen
			rightBorder = sepsIndices[i+1]
			result = append(result, string(srunes[leftBorder:rightBorder]))
		} else {
			leftBorder = sepsIndices[i] + seplen
			result = append(result, string(srunes[leftBorder:]))
		}
	}
	return result
}
