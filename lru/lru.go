package lru

import (
	"sync"
)

type node struct {
	next     *node
	previous *node
	key      *string
	content  [][]byte
}

func (n *node) delete() {
	n.next = nil
	n.previous = nil
	n.content = nil
	n.key = nil
}

// Concurrent safe LRU cache
type cache struct {
	*sync.Mutex

	indexes map[string]uint
	list    []*node

	head uint
	tail uint
	size uint
}
