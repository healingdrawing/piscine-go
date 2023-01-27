package piscine

func IsPrime(num int) bool {
	if num > 1 {
		if num == 2 || num == 3 {
			return true
		}
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}
