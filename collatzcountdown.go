package piscine

func CollatzCountdown(start int) int {
	// initiate step count
	if start < 1 {
		return -1
	}
	steps := 0
	for start > 1 {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = start*3 + 1
		}
		// steps collector
		steps++
	}
	return steps
}
