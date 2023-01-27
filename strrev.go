package piscine

func StrRev(s string) string {
	sar := []rune(s)
	for i, j := 0, len(sar)-1; i < j; i, j = i+1, j-1 {
		sar[i], sar[j] = sar[j], sar[i]
	}
	return string(sar[:])
}
