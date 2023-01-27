package piscine

func Fibonacci(index int) int {
	return firec(index)
}

func firec(index int) int {
	if index < 0 {
		return -1
	} else if index <= 1 {
		return index
	} else {
		return firec(index-1) + firec(index-2)
	}
}
