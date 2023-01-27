package main

import "github.com/01-edu/z01"

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

const (
	OPEN  = 1
	CLOSE = 0
)

type Door struct {
	state int
}

func OpenDoor(ptrDoor *Door) bool {
	printStr("Door Opening...")
	ptrDoor.state = OPEN
	return true
}

func CloseDoor(ptrDoor *Door) bool {
	printStr("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	printStr("is the Door opened ?")
	return ptrDoor.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	printStr("is the Door closed ?")
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}
	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
