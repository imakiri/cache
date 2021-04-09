package lru

import (
	"github.com/imakiri/erres"
	"sync"
)

type node struct {
	next     *node
	previous *node
	key      *string
	value    [][]byte
}

func (n *node) cutTail() {
	if n.previous != nil {
		panic("critical internal error: cache/lru/node.wipe() can be performed only on tail node")
	}
	if n.next != nil {
		n.next.previous = nil
	}
	n.next = nil
	n.value = nil
	n.key = nil
}

type Cache interface {
	Get(key string) ([][]byte, bool)
	Set(key string, value [][]byte) error
}

// Concurrent safe LRU cache
type cache struct {
	mutex *sync.Mutex

	indexes map[string]uint
	list    []*node
	head    uint
	tail    uint

	size           uint
	mutable        bool
	acceptNilValue bool
}

func (c *cache) Get(key string) ([][]byte, bool) {
	if key == "" {
		return nil, false
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	var index, ok = c.indexes[key]
	if !ok {
		return nil, false
	}

	var node = c.list[index]
	switch {
	case index == c.head:
		break
	case index == c.tail:
		if node.next != nil {
			c.tail = c.indexes[*node.next.key]
			node.next.previous = nil
		}
		node.previous = c.list[c.head]
		node.next = nil
		c.list[c.head].next = node
		c.head = index
	default:
		node.previous.next = node.next
		node.previous = c.list[c.head]
		node.next.previous = node.previous
		node.next = nil
		c.list[c.head].next = node
		c.head = index
	}

	return node.value, true
}

func (c *cache) Set(key string, value [][]byte) error {
	switch {
	case key == "":
		return erres.InvalidArgument.Extend(0).SetDescription("key string cannot be empty")
	case value == nil && !c.acceptNilValue:
		return erres.NilArgument.Extend(0).SetDescription("value cannot be nil")
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	var index, ok = c.indexes[key]
	switch {
	case ok && !c.mutable:
		return erres.AlreadyExist.Extend(0).SetDescription("given key already exist")
	case ok && c.mutable:
		c.list[index].value = value
		return nil
	case !ok:
		var node = new(node)
		node.key = &key
		node.value = value

		var size = uint(len(c.list))
		if size == c.size {
			var last = size - 1
			switch {
			case c.tail == last:
				c.list[last] = nil
				c.list = c.list[:c.tail]
			case c.tail < last:
				copy(c.list[c.tail:], c.list[c.tail+1:])
			}

			var tail = c.list[c.tail]
			c.tail = c.indexes[*tail.next.key]
			delete(c.indexes, *tail.key)
			tail.cutTail()
		}

		c.list = append(c.list, node)
		c.indexes[key] = size - 1

		switch {
		case index == 0:
			break
		default:
			var head = c.list[c.head]
			node.previous = head
			head.next = node
		}
	}

	return nil
}

func NewCache(size uint, mutable bool, acceptNilValue bool) (Cache, error) {
	if size == 0 {
		return nil, erres.InvalidArgument.Extend(0).SetDescription("size cannot be zero")
	}

	var cache = new(cache)

	cache.mutex = new(sync.Mutex)
	cache.indexes = make(map[string]uint, size)
	cache.list = make([]*node, 0, size)

	cache.size = size
	cache.mutable = mutable
	cache.acceptNilValue = acceptNilValue

	return cache, nil
}
