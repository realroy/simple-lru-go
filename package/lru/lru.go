package lru

// LRU ...
type LRU struct {
	Size    int
	MaxSize int
	ddl     *doubleLinkedList
	dict    map[string]*node
}

// NewLRU ...
func NewLRU(maxSize int) *LRU {
	return &LRU{
		Size:    0,
		MaxSize: maxSize,
		ddl:     newDoubleLinkedList(),
		dict:    map[string]*node{},
	}
}

// Set ...
func (lru *LRU) Set(key string, value interface{}) {
	if lru.Size+1 != lru.MaxSize {
		lru.Size++
	} else {
		lru.ddl.RemoveTail()
		delete(lru.dict, key)
	}

	node := lru.ddl.InsertToFront(key)
	lru.dict[key] = node
}

// Get ...
func (lru *LRU) Get(key string) interface{} {
	node, found := lru.dict[key]

	if !found {
		return nil
	}

	lru.ddl.MoveToFront(node)

	return node.Value
}
