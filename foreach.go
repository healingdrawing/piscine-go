package piscine

func ForEach(f func(int), a []int) {
	for _, ai := range a {
		f(ai)
	}
}
