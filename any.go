package piscine

func Any(f func(string) bool, a []string) bool {
	for _, ai := range a {
		if f(ai) {
			return true
		}
	}
	return false
}
