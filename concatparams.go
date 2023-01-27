package piscine

func ConcatParams(args []string) string {
	lena := len(args) - 1
	sepa := "\n"
	var sums string
	for ind, arg := range args {
		sums += arg
		if ind < lena {
			sums += sepa
		}
	}
	return sums
}
