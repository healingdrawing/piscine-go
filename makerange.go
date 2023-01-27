package piscine

func MakeRange(min, max int) []int {
	var arr []int
	arr = nil
	if min < max {
		arr = make([]int, max-min)
		for i := min; i < max; i++ {
			arr[i-min] = i
		}
		return arr
	}
	return arr
}
