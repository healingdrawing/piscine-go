package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 != nil && n2 != nil {
		makestep := true
		for makestep {
			n1 = SortListInsert(n1, n2.Data)
			if n2.Next != nil {
				n2 = n2.Next
			} else {
				makestep = false
			}
		}
	}
	return n1
}
