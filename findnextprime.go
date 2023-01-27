package piscine

func FindNextPrime(nb int) int {
	number := -1
	for i := nb; ; i++ {
		if IsPrimeRapid(i) {
			number = i
			break
		}
	}
	return number
}

func IsPrimeRapid(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1
}

// func IsPrimeRapid(num int) bool {
// 	if num > 1 {
// 		if num == 2 || num == 3 {
// 			return true
// 		}
// 		for i := 2; i < num; i++ {
// 			if num%i == 0 {
// 				return false
// 			}
// 		}
// 		return true
// 	}
// 	return false
// }
