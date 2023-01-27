package piscine

// func main() {
// 	fmt.Println(BasicAtoi("120"))
// }

func BasicAtoi(s string) int {
	sar := []rune(s)
	ln := len(sar)
	numsum := 0
	for i := 0; i < ln; i++ {
		numsum = numsum*10 + (int(sar[i] - '0')) // * powerNu(10, ln-i-1)
	}
	return numsum
}

// func BasicAtoiX(s string) int {
// 	sar := []rune(s)
// 	l := len(sar)
// 	numsum := 0
// 	for i := len(sar) - 1; i > -1; i-- {
// 		numsum += int(sar[i]) * powerNu(10, l-i)
// 	}
// 	return numsum
// }

// func powerNu(n int, power int) (r int) {
// 	r = n
// 	for i := 2; i < power; i++ {
// 		r *= n
// 	}
// 	return
// }
