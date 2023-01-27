package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	template := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	digits := PositiveIntToDigits(n)
	digits = SortIntArrAscending(digits)
	for _, r := range digits {
		z01.PrintRune(template[r])
	}
}

func PositiveIntToDigits(n int) (res []int) {
	if n > 9 {
		for n > 9 {
			tail := n % 10
			n = n / 10
			res = append(res, tail)
		}
		res = append(res, n)
		for i2, j := 0, len(res)-1; i2 < j; i2, j = i2+1, j-1 {
			res[i2], res[j] = res[j], res[i2]
		}

	} else {
		res = append(res, n)
	}
	return
}

func SortIntArrAscending(aon []int) []int {
	lena := len(aon) - 1
	sortedPairs := 0
	for sortedPairs < lena {
		for i := 0; i < lena; i++ {
			if aon[i] > aon[i+1] {
				sortedPairs = 0
				aon[i], aon[i+1] = aon[i+1], aon[i]
			} else {
				sortedPairs++
			}
		}
	}
	return aon
}
