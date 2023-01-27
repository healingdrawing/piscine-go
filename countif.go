package piscine

func CountIf(f func(string) bool, tab []string) int {
	result := 0
	for _, tabi := range tab {
		if f(tabi) {
			result++
		}
	}
	return result
}
