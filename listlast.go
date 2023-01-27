package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListLast(l *List) interface{} {
	// check the nodes
	if l.Head == nil {
		return nil
	}
	n := l.Head
	for n != nil {
		if n.Next == nil {
			return n.Data
		}
		n = n.Next
	}
	return nil
}
