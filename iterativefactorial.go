package piscine

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

func IterativeFactorial(nb int) int {
	result := 0

	if nb == 0 {
		return 1
	} else if nb < 0 {
		return 0
	} else if nb > 0 && nb < 21 {
		fac := 1
		for i := nb; i > 1; i-- {
			fac *= i
		}
		return fac
	}

	return result
}
