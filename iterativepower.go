package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	result := nb
	for i := power; i > 1; i-- {
		result *= nb
	}
	return result
}
