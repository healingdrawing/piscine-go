package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n == 1 {
		PrintCombN1()
	} else if n == 2 {
		PrintCombN2()
	} else if n == 3 {
		PrintCombN3()
	} else if n == 4 {
		PrintCombN4()
	} else if n == 5 {
		PrintCombN5()
	} else if n == 6 {
		PrintCombN6()
	} else if n == 7 {
		PrintCombN7()
	} else if n == 8 {
		PrintCombN8()
	} else if n == 9 {
		PrintCombN9()
	}
}

func PrintCombN1() {
	printail := true
	for a := 0; a < 10; a++ {
		if a == 0 {
			printail = false
		}
		if printail {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		z01.PrintRune(rune(IRn(a)))
		printail = true
	}
	z01.PrintRune('\n')
}

func PrintCombN2() {
	printail := true
	for a := 0; a < 10; a++ {
		for b := a; b < 10; b++ {
			if a < b {
				if a == 0 && b == 1 {
					printail = false
				}
				if printail {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				z01.PrintRune(rune(IRn(a)))
				z01.PrintRune(rune(IRn(b)))
				printail = true
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintCombN3() {
	printail := true
	for a := 0; a < 10; a++ {
		for b := a; b < 10; b++ {
			for c := b; c < 10; c++ {
				if a < b && b < c {
					if a == 0 && b == 1 && c == 2 {
						printail = false
					}
					if printail {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(rune(IRn(a)))
					z01.PrintRune(rune(IRn(b)))
					z01.PrintRune(rune(IRn(c)))
					printail = true
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintCombN4() {
	printail := true
	for a := 0; a < 10; a++ {
		for b := a; b < 10; b++ {
			for c := b; c < 10; c++ {
				for d := c; d < 10; d++ {
					if a < b && b < c && c < d {
						if a == 0 && b == 1 && c == 2 && d == 3 {
							printail = false
						}
						if printail {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
						z01.PrintRune(rune(IRn(a)))
						z01.PrintRune(rune(IRn(b)))
						z01.PrintRune(rune(IRn(c)))
						z01.PrintRune(rune(IRn(d)))
						printail = true
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintCombN5() {
	printail := true
	for a := 0; a < 10; a++ {
		for b := a; b < 10; b++ {
			for c := b; c < 10; c++ {
				for d := c; d < 10; d++ {
					for e := d; e < 10; e++ {
						if a < b && b < c && c < d && d < e {
							if a == 0 && b == 1 && c == 2 && d == 3 && e == 4 {
								printail = false
							}
							if printail {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
							z01.PrintRune(rune(IRn(a)))
							z01.PrintRune(rune(IRn(b)))
							z01.PrintRune(rune(IRn(c)))
							z01.PrintRune(rune(IRn(d)))
							z01.PrintRune(rune(IRn(e)))
							printail = true
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintCombN6() {
	printail := true
	for a := 0; a < 10; a++ {
		for b := a; b < 10; b++ {
			for c := b; c < 10; c++ {
				for d := c; d < 10; d++ {
					for e := d; e < 10; e++ {
						for f := e; f < 10; f++ {
							if a < b && b < c && c < d && d < e && e < f {
								if a == 0 && b == 1 && c == 2 && d == 3 && e == 4 && f == 5 {
									printail = false
								}
								if printail {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
								z01.PrintRune(rune(IRn(a)))
								z01.PrintRune(rune(IRn(b)))
								z01.PrintRune(rune(IRn(c)))
								z01.PrintRune(rune(IRn(d)))
								z01.PrintRune(rune(IRn(e)))
								z01.PrintRune(rune(IRn(f)))
								printail = true
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintCombN7() {
	printail := true
	for a := 0; a < 10; a++ {
		for b := a; b < 10; b++ {
			for c := b; c < 10; c++ {
				for d := c; d < 10; d++ {
					for e := d; e < 10; e++ {
						for f := e; f < 10; f++ {
							for x := f; x < 10; x++ {
								if a < b && b < c && c < d && d < e && e < f && f < x {
									if a == 0 && b == 1 && c == 2 && d == 3 && e == 4 && f == 5 && x == 6 {
										printail = false
									}
									if printail {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
									z01.PrintRune(rune(IRn(a)))
									z01.PrintRune(rune(IRn(b)))
									z01.PrintRune(rune(IRn(c)))
									z01.PrintRune(rune(IRn(d)))
									z01.PrintRune(rune(IRn(e)))
									z01.PrintRune(rune(IRn(f)))
									z01.PrintRune(rune(IRn(x)))
									printail = true
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintCombN8() {
	printail := true
	for a := 0; a < 10; a++ {
		for b := a; b < 10; b++ {
			for c := b; c < 10; c++ {
				for d := c; d < 10; d++ {
					for e := d; e < 10; e++ {
						for f := e; f < 10; f++ {
							for x := f; x < 10; x++ {
								for y := x; y < 10; y++ {
									if a < b && b < c && c < d && d < e && e < f && f < x && x < y {
										if a == 0 && b == 1 && c == 2 && d == 3 && e == 4 && f == 5 && x == 6 && y == 7 {
											printail = false
										}
										if printail {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
										z01.PrintRune(rune(IRn(a)))
										z01.PrintRune(rune(IRn(b)))
										z01.PrintRune(rune(IRn(c)))
										z01.PrintRune(rune(IRn(d)))
										z01.PrintRune(rune(IRn(e)))
										z01.PrintRune(rune(IRn(f)))
										z01.PrintRune(rune(IRn(x)))
										z01.PrintRune(rune(IRn(y)))
										printail = true
									}
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintCombN9() {
	printail := true
	for a := 0; a < 10; a++ {
		for b := a; b < 10; b++ {
			for c := b; c < 10; c++ {
				for d := c; d < 10; d++ {
					for e := d; e < 10; e++ {
						for f := e; f < 10; f++ {
							for x := f; x < 10; x++ {
								for y := x; y < 10; y++ {
									for z := y; z < 10; z++ {
										if a < b && b < c && c < d && d < e && e < f && f < x && x < y && y < z {
											if a == 0 && b == 1 && c == 2 && d == 3 && e == 4 && f == 5 && x == 6 && y == 7 && z == 8 {
												printail = false
											}
											if printail {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
											z01.PrintRune(rune(IRn(a)))
											z01.PrintRune(rune(IRn(b)))
											z01.PrintRune(rune(IRn(c)))
											z01.PrintRune(rune(IRn(d)))
											z01.PrintRune(rune(IRn(e)))
											z01.PrintRune(rune(IRn(f)))
											z01.PrintRune(rune(IRn(x)))
											z01.PrintRune(rune(IRn(y)))
											z01.PrintRune(rune(IRn(z)))
											printail = true
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func IRn(s int) (r rune) {
	if s == 0 {
		r = '0'
	} else if s == 1 {
		r = '1'
	} else if s == 2 {
		r = '2'
	} else if s == 3 {
		r = '3'
	} else if s == 4 {
		r = '4'
	} else if s == 5 {
		r = '5'
	} else if s == 6 {
		r = '6'
	} else if s == 7 {
		r = '7'
	} else if s == 8 {
		r = '8'
	} else if s == 9 {
		r = '9'
	}
	return
}
