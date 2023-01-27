package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListReverse(l *List) {
	// var prev *NodeL = nil
	// current := l.Head
	// for current != nil {
	// 	next := current.Next
	// 	current.Next = prev
	// 	prev = current
	// 	current = next
	// }
	// l.Head = prev

	var prev *NodeL = nil
	var current *NodeL = l.Head
	var next *NodeL = nil
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
