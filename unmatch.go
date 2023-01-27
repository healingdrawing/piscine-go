package piscine

func Unmatch(a []int) int {
	var varCount []bool
	for i := 0; i < len(a); i++ {
		varCount = append(varCount, false)
	}
	// once fired pair of indices become true
	for i := range a {
		for j := range a {
			if a[i] == a[j] && i != j && !varCount[i] && !varCount[j] {
				varCount[i] = true
				varCount[j] = true
			}
		}
	}
	for i, p := range varCount {
		if !p {
			return a[i]
		}
	}
	return -1
}
