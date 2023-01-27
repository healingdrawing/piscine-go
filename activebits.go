package piscine

// print number of 1 in binary form of incoming integer
func ActiveBits(n int) int {
	intBits := itoi01(n)
	return counterOf1(intBits)
}

// integer to binary array of int 0 1
// convert integer number to int array of binary form representation
// right(first) byte in index 0
func itoi01(n int) []int {
	var ar []int
	for n > 0 {
		ar = append(ar, n%2)
		n /= 2
	}
	return ar
}

// counter 1 inside int array
func counterOf1(intBits []int) int {
	counter := 0
	for _, item := range intBits {
		if item > 0 {
			counter++
		}
	}
	return counter
}
