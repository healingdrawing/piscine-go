package piscine

func Join(strs []string, sep string) string {
	result := ""
	lena := len(strs)
	for i, v := range strs {
		result += v
		if i < lena-1 {
			result += sep
		}
	}
	return result
}
