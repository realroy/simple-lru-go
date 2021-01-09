package lru

// Node ...
type node struct {
	Next  *node
	Prev  *node
	Value interface{}
}

// newNode ...
func newNode(next *node, prev *node, value interface{}) *node {
	return &node{
		Next:  next,
		Prev:  prev,
		Value: value,
	}
}
