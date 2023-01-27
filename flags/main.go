package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	flagInsert := false
	flagInsertArgInd := -1 // index of flag inside args array
	flagOrder := false
	flagOrderArgInd := -1 // index of flag inside args array

	// check the flags
	for ind, arg := range args {
		if !flagInsert {
			indInsert := sISInd("--insert=", arg)
			if indInsert > -1 { // if flag found write his index
				flagInsertArgInd = ind
			}
			if indInsert == -1 {
				indInsert = sISInd("-i=", arg)
				if indInsert > -1 { // if flag found write his index
					flagInsertArgInd = ind
				}
			}
			if indInsert > -1 {
				flagInsert = true
			}
		}

		if !flagOrder {
			indOrder := sISInd("--order", arg)
			if indOrder > -1 { // if flag found write his index
				flagOrderArgInd = ind
			}
			if indOrder == -1 {
				indOrder = sISInd("-o", arg)
				if indOrder > -1 { // if flag found write his index
					flagOrderArgInd = ind
				}
			}
			if indOrder > -1 {
				flagOrder = true
			}
		}

		// fmt.Println("ind = ", ind, "  arg = ", arg)
	}
	// calculation of the output

	if len(args) > 1 && args[1] == "-h" ||
		len(args) > 1 && args[1] == "--help" ||
		len(args) == 1 { // help print
		printHelp()
	} else if len(args) > 1 && args[1] != "-h" ||
		len(args) > 1 && args[1] != "--help" { // usual print
		Coto(args, flagInsert, flagOrder, flagInsertArgInd, flagOrderArgInd)
	}
}

func printHelp() {
	help := "--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument."
	fmt.Println(help)
}

func Coto(args []string, flagInsert bool, flagOrder bool, flagInsertArgInd int, flagOrderArgInd int) {
	var fpRunes []rune           // final printable runes
	if flagInsert && flagOrder { // order output
		// collecting runes sequence
		// --insert= section
		insertArg := args[flagInsertArgInd]
		equalIndex := sISInd("=", insertArg)
		insertRunes := []rune(args[flagInsertArgInd])
		ipRunes := insertRunes[equalIndex+1:] // insert printable runes, slice from = to end
		// other section
		// filter flags
		var notFlagsData []string
		for i := 1; i < len(args); i++ { // zero ignored as name of programm
			if i != flagInsertArgInd && i != flagOrderArgInd {
				notFlagsData = append(notFlagsData, args[i])
			}
		}
		// collect runes from all flags
		var nfRunes []rune
		for _, sbox := range notFlagsData {
			runes := []rune(sbox)
			nfRunes = append(nfRunes, runes...)
		}
		fpRunes = append(fpRunes, ipRunes...)
		fpRunes = append(fpRunes, nfRunes...)
		fpRunes = sortRunesAscending(fpRunes)
	} else if flagInsert && !flagOrder { // add insert text at the end as in example
		// collecting runes sequence
		// --insert= section
		insertArg := args[flagInsertArgInd]
		equalIndex := sISInd("=", insertArg)
		insertRunes := []rune(args[flagInsertArgInd])
		ipRunes := insertRunes[equalIndex+1:] // insert printable runes, slice from = to end
		// other section
		// filter flags
		var notFlagsData []string
		for i := 1; i < len(args); i++ { // zero ignored as name of programm
			if i != flagInsertArgInd && i != flagOrderArgInd {
				notFlagsData = append(notFlagsData, args[i])
			}
		}
		// collect runes from all flags
		var nfRunes []rune
		for _, sbox := range notFlagsData {
			runes := []rune(sbox)
			nfRunes = append(nfRunes, runes...)
		}
		fpRunes = append(fpRunes, nfRunes...)
		fpRunes = append(fpRunes, ipRunes...)
	} else if !flagInsert && !flagOrder { // just print incoming params
		// collecting runes sequence
		// --insert= section

		// insertArg := args[flagInsertArgInd]
		// equalIndex := sISInd("=", insertArg)
		// insertRunes := []rune(args[flagInsertArgInd])
		// ipRunes := insertRunes[equalIndex+1:] // insert printable runes, slice from = to end

		// other section
		// filter flags
		var notFlagsData []string
		for i := 1; i < len(args); i++ { // zero ignored as name of programm
			if i != flagInsertArgInd && i != flagOrderArgInd {
				notFlagsData = append(notFlagsData, args[i])
			}
		}
		// collect runes from all flags
		var nfRunes []rune
		for _, sbox := range notFlagsData {
			runes := []rune(sbox)
			nfRunes = append(nfRunes, runes...)
		}
		// fpRunes = append(fpRunes, ipRunes...)
		fpRunes = append(fpRunes, nfRunes...)
		// fpRunes = sortRunesAscending(fpRunes)
	} else if !flagInsert && flagOrder {
		// collecting runes sequence
		// --insert= section

		// insertArg := args[flagInsertArgInd]
		// equalIndex := sISInd("=", insertArg)
		// insertRunes := []rune(args[flagInsertArgInd])
		// ipRunes := insertRunes[equalIndex+1:] // insert printable runes, slice from = to end

		// other section
		// filter flags
		var notFlagsData []string
		for i := 1; i < len(args); i++ { // zero ignored as name of programm
			if i != flagInsertArgInd && i != flagOrderArgInd {
				notFlagsData = append(notFlagsData, args[i])
			}
		}
		// collect runes from all flags
		var nfRunes []rune
		for _, sbox := range notFlagsData {
			runes := []rune(sbox)
			nfRunes = append(nfRunes, runes...)
		}
		// fpRunes = append(fpRunes, ipRunes...)
		fpRunes = append(fpRunes, nfRunes...)
		fpRunes = sortRunesAscending(fpRunes)
	}

	runesPrinter(fpRunes)
}

func runesPrinter(runes []rune) {
	fmt.Println(string(runes))
}

func sortRunesAscending(runes []rune) []rune {
	lena := len(runes) - 1
	sortedPairs := 0
	for sortedPairs < lena {
		for i := 0; i < lena; i++ {
			if runes[i] > runes[i+1] {
				sortedPairs = 0
				runes[i], runes[i+1] = runes[i+1], runes[i]
			} else {
				sortedPairs++
			}
		}
	}
	return runes
}

// Returns index of substring s1 inside string s2 or -1. Remastered from GeeksForGeeks
func sISInd(s1 string, s2 string) int {
	r1 := []rune(s1)
	r2 := []rune(s2)

	M := len(r1)
	N := len(r2)

	/* A loop to slide pat[] one by one */
	for i := 0; i <= N-M; i++ {
		var j int

		/* For current index i, check for
		pattern match */
		for j = 0; j < M; j++ {
			if s2[i+j] != s1[j] {
				break
			}
		}
		if j == M {
			return i
		}
	}
	return -1
}

func pL(s string) { fmt.Println(s) }

// search string inside other string
// func stringInsideString(whatCheck string, whereCheck string) bool {
// 	result := false
// 	if len(whatCheck) > len(whereCheck) {
// 		return false
// 	} else if whatCheck == whereCheck {
// 		return true
// 	} else { // string checking

// 		howManyFound := 0 // how many runes found in row(sequence)
// 		whereLastInd := 0 // last index of iteration of whereCheck
// 		whatLastInd := 0  // last index of rune used for search

// 		whatRunes := []rune(whatCheck)
// 		whatLen := len(whatCheck) // length of checked sequence

// 		whereRunes := []rune(whereCheck)
// 		whereLen := len(whereRunes)

// 		for howManyFound < whatLen || whereLastInd > whereLen-whatLen {
// 			whatRune := whatRunes[whatLastInd] // what need to check

// 			runeInside, indInsideWhere := RuneInsideString(whatRune, whereCheck, whereLastInd)
// 			if runeInside {
// 				whatLastInd++
// 				howManyFound++
// 				whereLastInd = indInsideWhere
// 				result = true
// 			} else {
// 				whatLastInd = 0
// 				howManyFound = 0
// 				whereLastInd -= howManyFound
// 				result = false
// 			}
// 		}
// 	}
// 	return result
// }

// //
// func RuneInsideString(r rune, s string, startInd int) (runeinside bool, index int) {
// 	runeinside = false
// 	index = -1
// 	for ind, srune := range s {
// 		if ind > startInd {
// 			if srune == r {
// 				runeinside = true
// 				index = ind
// 			}
// 		}
// 	}
// 	return
// }
