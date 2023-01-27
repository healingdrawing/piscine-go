package piscine

func Compare(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)
	lena := len(a)
	lenb := len(b)
	var lenc int
	if lena > lenb {
		lenc = lenb
	} else {
		lenc = lena
	}
	for i := 0; i < lenc; i++ {
		if ra[i] > rb[i] {
			return 1
		} else if ra[i] < rb[i] {
			return -1
		}
	}
	if lena == lenb {
		return 0
	} else if lena > lenb {
		return 1
	} else {
		return -1
	}
}
