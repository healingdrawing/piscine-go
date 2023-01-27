package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	current := l // copy to keep link to list beginning in original form
	if current != nil && current.Next == nil {
		return nil
	} else {
		for current != nil {
			index := current.Next
			for index != nil {
				if current.Data > index.Data {
					current.Data, index.Data = index.Data, current.Data
				}
				index = index.Next
			}
			current = current.Next // if not used current(but l) then it will cut the list
		}
	}
	return l
}
