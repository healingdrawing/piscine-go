package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListRemoveIf(l *List, data_ref interface{}) {
	// Store head node
	temp := l.Head
	var prev *NodeL = nil

	// If head node itself holds the key
	// or multiple occurrences of key
	for temp != nil && temp.Data == data_ref {
		l.Head = temp.Next // Changed head
		temp = l.Head      // Change Temp
	}

	// Delete occurrences other than head
	for temp != nil {

		// Search for the key to be deleted,
		// keep track of the previous node
		// as we need to change 'prev.next'
		for temp != nil && temp.Data != data_ref {
			prev = temp
			temp = temp.Next
		}

		// If key was not present in linked list
		if temp == nil {
			return // l.Head
		}

		// Unlink the node from linked list
		prev.Next = temp.Next

		// Update Temp for next iteration of outer loop
		temp = prev.Next
	}
}

// if self.head is None:
// return
// if position == 0:
// self.head = self.head.next
// return self.head
// index = 0
// current = self.head
// prev = self.head
// temp = self.head
// while current is not None:
// if index == position:
// 	temp = current.next
// 	break
// prev = current
// current = current.next
// index += 1
// prev.next = temp
// return prev
