package piscine

func Abort(a, b, c, d, e int) int {
	data := []int{a, b, c, d, e}
	// sorting ascending
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	// median calculation
	var median int
	l := len(data)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (data[l/2-1] + data[l/2]) / 2
	} else {
		median = data[l/2]
	}
	return median
}

// https://www.includehelp.com/c-programs/calculate-median-of-an-array.aspx
