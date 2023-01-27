package piscine

// func main() {
// 	fmt.Println(BasicAtoi2("1209ASD"))
// }

func BasicAtoi2(s string) int {
	sign := 1
	startindex := 0
	var num int

	if s == "" || s == "0" {
		return 0
	}
	if s[0] == '-' {
		startindex = 1
		sign = -1
	}
	if s[0] == '+' {
		startindex = 1
	}

	for i := startindex; i < len(s); i++ {
		if rune(s[1]) < '0' || rune(s[1]) > '9' {
			return 0
		}
		num = num*10 + int(rune(s[i])-'0')
	}
	return num * sign
}
