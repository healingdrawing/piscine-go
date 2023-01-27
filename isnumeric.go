package piscine

func IsNumeric(s string) bool {
	template = []rune{
		'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
	}
	runes := []rune(s)
	for _, r := range runes {
		result := false

		for _, tr := range template {
			if tr == r {
				result = true
			}
		}
		if result == false {
			return false
		}
	}
	return true
}
