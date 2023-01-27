package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, ai := range a {
		result = append(result, f(ai))
	}
	return result
}
