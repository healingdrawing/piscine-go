package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

// func ListSize(l *List) int {
// 	var size int
// 	if l == nil {
// 		size = 0
// 	} else {
// 		size = 1
// 		last := l.Head
// 		for last.Next != nil {
// 			last = last.Next
// 			size++
// 		}
// 	}
// 	return size
// }

func ListSize(l *List) int {
	var temp *NodeL = l.Head
	var count int = 0
	for temp != nil {
		// Visit to next node
		temp = temp.Next
		// Increase the node counter
		count++
	}
	return count
}
