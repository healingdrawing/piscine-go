package piscine

func Compact(ptr *[]string) int {
	counter := 0
	var newSlice []string
	for _, item := range *ptr {
		if item != "" {
			newSlice = append(newSlice, item+"")
			counter++
		}
	}
	*ptr = newSlice
	return counter
}
