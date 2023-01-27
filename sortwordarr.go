package piscine

// func main() {
// 	a := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
// 	SortWordArr(a)
// }

func SortWordArr(a []string) {
	pairs := len(a) - 1 // pairs in array
	sortedPairs := 0    // completely sorted at the moment
	for sortedPairs < pairs {
		for i := 0; i < pairs; i++ {
			// in task the ascending sorting is required, so 1 is bad response
			if compare(a[i], a[i+1]) == 1 {
				a[i], a[i+1] = a[i+1], a[i]
				sortedPairs = 0
			} else {
				sortedPairs++
			}
		}
	}
	// some magic print required and in same time allowed only built in functionality
	// fmt.Println(a)
}

// (descending) if a > b return 1. (ascending) if a > b return -1. zero if equal
func compare(a, b string) int {
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
