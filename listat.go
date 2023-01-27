package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

func ListAt(l *NodeL, pos int) *NodeL {
	current := l
	count := 0
	/*
	   index of Node we are currently looking at
	*/
	for current != nil {
		if count == pos {
			return current
		}
		count++
		current = current.Next
	}
	return nil
}

// https://www.geeksforgeeks.org/data-structures/linked-list/#singlyLinkedList
