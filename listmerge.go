package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListMerge(l1 *List, l2 *List) {
	n := l2.Head
	for n != nil {
		ListPushBack(l1, n.Data)
		n = n.Next
	}
}
