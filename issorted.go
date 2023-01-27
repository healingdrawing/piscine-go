package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	fresh := true      // fresh state, no choosen sorting direction
	ascending := false // false used too as descending
	for i := 0; i < len(a)-1; i++ {
		if fresh { // choose the sorting direction if elements not equal
			if f(a[i], a[i+1]) != 0 {
				ascending = a[i] < a[i+1]
				fresh = false
			}
		} else { // sorting direction is already choosen
			if ascending && f(a[i], a[i+1]) > 0 ||
				!ascending && f(a[i], a[i+1]) < 0 {
				return false
			}
		}
	}
	return true
}
