package main

import (
	"os"
)

func main() {
	// tn, _ := multiplyExact(int64(-1), int64(1))
	// osPrinter(itos(int(tn)))
	// fmt.Println(int64(459) / int64(860))
	osargs := os.Args
	var args []string
	// check arguments number
	if len(osargs) != 4 {
		return
	} else {
		for i := 1; i < len(osargs); i++ {
			args = append(args, osargs[i])
		}

		// osPrinter("---after loop---- len agrs = " + itos(len(args)))
	}
	// raw arguments
	rawFirst := args[0]
	rawSecond := args[1] // raw operator
	rawThird := args[2]
	// check the operator
	operator := checkTheOperator(rawSecond)
	if operator == 0 {
		return // nothing
	}
	// check incoming numbers
	// osPrinter("\nrawFirst == " + rawFirst)
	good1, n1 := checkNumber(rawFirst)
	// osPrinter("\nn1 == " + itos(n1))
	if !good1 { // not a number
		return
	}

	// osPrinter("\nrawThird == " + rawThird)
	good2, n2 := checkNumber(rawThird)
	// osPrinter("\nn2 == " + itos(n2))
	if !good2 {
		return
	}
	// check the required division rules not broken
	if n2 == 0 && operator == 4 {
		osPrinter("No division by 0" + "\n")
		return
	} else if n2 == 0 && operator == 5 {
		osPrinter("No modulo by 0" + "\n")
		return
	}
	result, status := calculation(n1, operator, n2)
	if status != "ok" {
		return
	} else {
		osPrinter(itos(result) + "\n")
	}
}

func osPrinter(s string) {
	os.Stdout.WriteString(s)
}

// make calc
func calculation(n1, operator, n2 int) (result int, status string) {
	var result64 int64
	status = "ok"
	if operator == 1 {
		result64 = int64(n1) + int64(n2)
	} else if operator == 2 {
		result64 = int64(n1) - int64(n2)
	} else if operator == 3 {
		result64 = int64(n1) * int64(n2)
	} else if operator == 4 {
		// osPrinter("fired division " + itos(n1) + "/" + itos(n2))
		result64 = int64(n1) / int64(n2)
		// os.Stdout.WriteString("inside calc " + strconv.Itoa(int(result64)))
	} else {
		result64 = int64(n1) % int64(n2)
	}
	result = int(result64)
	// os.Stdout.WriteString("inside calc result " + strconv.Itoa(result))
	// os.Stdout.WriteString("inside calc result2 " + itos(result))
	// check to overflow happened
	if isOverflowHappened(n1, n2, operator, result64) {
		status = "overflow"
	}
	// osPrinter("\nstatus = " + status)
	// osPrinter("\nresult = " + itos(result))
	return
}

// func lessOfTwoHasModulo(n1 int, n2 int) bool {
// 	if n1 < n2 && n1%2 > 0 {
// 		return true
// 	} else if n2 > n1 && n2%2 > 0 {
// 		return true
// 	}

// 	// if n1%2 == 1 && n2%2 == 0 || n1%2 == 0 && n2%2 == 1 {
// 	// 	return true
// 	// }
// 	return false
// }

const (
	mostNegative = -(mostPositive + 1)
	mostPositive = 1<<63 - 1
)

func multiplyExact(a, b int64) (int64, bool) {
	// osPrinter("\nmulti fired " + itos(int(a)) + " " + itos(int(b)))
	result := a * b
	if a == 0 || b == 0 || a == 1 || b == 1 {
		return result, false
	}
	if a == mostNegative || b == mostNegative {
		return result, true
	}
	if result/b != a {
		return result, true
	}
	return result, false
}

func isOverflowHappened(n1, n2, operator int, result64 int64) bool {
	if operator == 1 { // +
		if n1 > 0 && n2 > 0 && result64 < 0 {
			return true
		} else if n1 < 0 && n2 < 0 && result64 > 0 {
			return true
		}
	} else if operator == 2 { // -
		if n1 > 0 && n2 < 0 && result64 < 0 ||
			n1 < 0 && n2 > 0 && result64 > 0 {
			return true
		}
	} else if operator == 3 { // *
		_, err := multiplyExact(int64(n1), int64(n2))
		if err {
			return true
		}
		// if n1 > 0 && n2 > 0 && result64 < 0 && !lessOfTwoHasModulo(n1, n2) ||
		// 	n1 < 0 && n2 < 0 && result64 < 0 && !lessOfTwoHasModulo(n1, n2) ||
		// 	n1 > 0 && n2 < 0 && result64 > 0 && !lessOfTwoHasModulo(n1, n2) ||
		// 	n1 < 0 && n2 > 0 && result64 > 0 && !lessOfTwoHasModulo(n1, n2) ||

		// 	n1 > 0 && n2 > 0 && result64 > 0 && lessOfTwoHasModulo(n1, n2) ||
		// 	n1 < 0 && n2 < 0 && result64 > 0 && lessOfTwoHasModulo(n1, n2) ||
		// 	n1 > 0 && n2 < 0 && result64 < 0 && lessOfTwoHasModulo(n1, n2) ||
		// 	n1 < 0 && n2 > 0 && result64 < 0 && lessOfTwoHasModulo(n1, n2) {
		// 	return true
		// }
	} else if operator == 4 { // /
		return false // at least no need check for int / int case
	} else if operator == 5 { // %
		return false // at least no need check for int % int case
	}
	return false
}

// string number to integer
func stoi(s string) int {
	sar := []rune(s)
	minus := false
	if sar[0] == '-' {
		minus = true
		sar = sar[1:]
	}
	ln := len(sar)
	su := 0
	for i := 0; i < ln; i++ {
		su = su*10 + (int(sar[i] - '0'))
	}
	if minus {
		su *= -1
	}
	return su
}

// integer number to string
func itos(n int) string {
	minus := n < 0
	te := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9"}
	var digits []int
	if n != 0 {
		for n != 0 {
			digits = append(digits, n%10)
			n /= 10
		}
	} else {
		return "0"
	}
	// reverse digits to correct way
	if len(digits) > 1 {
		for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
			digits[i], digits[j] = digits[j], digits[i]
		}
	}
	// collect string from digits
	s := ""
	for _, d := range digits {
		if d < 0 {
			s += te[-d]
		} else {
			s += te[d]
		}
		// osPrinter("\ninisde itos " + s)
	}
	// add minus
	if minus {
		s = "-" + s
	}
	return s
}

// only decimal digits inside
func onlyDecimalDigits(runes []rune) bool {
	for _, r := range runes {
		if r < '0' && r > '9' {
			return false
		}
	}
	return true
}

// if incoming number string form overflows after converting into int
func isIncomingNumberOverflow(s string) bool {
	// string to int number
	stinum := stoi(s)
	itsnum := itos(stinum)
	return s != itsnum
}

// check the number string form looks correct
func checkNumber(s string) (bool, int) {
	minus := false
	runes := []rune(s)
	if runes[0] == '-' {
		runes = runes[1:]
		minus = true
	}
	if !onlyDecimalDigits(runes) {
		return false, 0
	}
	// check incoming to overflow
	if isIncomingNumberOverflow(s) {
		return false, 0
	}
	// convert to integer
	su := stoi(string(runes))
	// check the negative is required
	if minus {
		su *= -1
	}
	return true, su
}

// check the incoming is correct operator and return index or 0 if wrong input data
func checkTheOperator(raw string) int {
	switch raw {
	case "+":
		return 1
	case "-":
		return 2
	case "*":
		return 3
	case "/":
		return 4
	case "%":
		return 5
	}
	return 0
}
