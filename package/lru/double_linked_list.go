package lru

// doubleLinkedList ...
type doubleLinkedList struct {
	Root *node
	Size int
}

// newdoubleLinkedList ...
func newDoubleLinkedList() *doubleLinkedList {
	return &doubleLinkedList{
		Root: newNode(nil, nil, nil),
		Size: 0,
	}
}

// MoveToFront ...
func (d *doubleLinkedList) MoveToFront(n0 *node) *node {
	n1 := n0.Next
	nMinusOne := n0.Prev

	if n1 != nil && nMinusOne != nil {
		nMinusOne.Next = n1
		n1.Prev = nMinusOne
	}

	oldFrontNode := d.Root.Next
	n0.Prev = d.Root
	n0.Next = oldFrontNode

	d.Root.Next = n0
	oldFrontNode.Prev = n0

	return n0
}

func (d *doubleLinkedList) InsertToFront(value interface{}) *node {
	d.Size++

	n := newNode(nil, nil, value)
	d.MoveToFront(n)

	return n
}

func (d *doubleLinkedList) RemoveTail() *node {
	node := d.Root.Prev
	prevNode := node.Prev

	node.Next = nil
	node.Prev = nil

	prevNode.Next = d.Root
	d.Root.Prev = prevNode

	d.Size--

	return node
}
