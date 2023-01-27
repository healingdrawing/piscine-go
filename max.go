package piscine

func Max(a []int) int {
	var max int
	if len(a) < 1 {
		return 0
	} else {
		max = -2147483648 // math.MinInt32
		for _, n := range a {
			if n > max {
				max = n
			}
		}
	}
	return max
}
