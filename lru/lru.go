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

func (n *node) drop() {
	if n.previous != nil {
		panic("cache/lru/node.drop() can be performed only on tail node")
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

type cache struct {
	mutex *sync.Mutex

	indexes map[string]*node
	head    *node
	tail    *node

	len            uint
	size           uint
	mutable        bool
	acceptNilValue bool
}

func (c *cache) build(node *node) {
	if node == nil {
		panic("cache/lru/node.build()  node cannot be nil")
	}

	switch {
	case c.len == 0:
		c.tail = node
		c.head = node
	default:
		c.head.next = node
		node.previous = c.head

		c.head = node
	}

	c.len++
}

func (c *cache) rebuild(node *node) {
	if node == nil {
		panic("cache/lru/node.rebuild()  node cannot be nil")
	}

	switch {
	case node == c.head:
		break
	case node == c.tail:
		if node.next == nil {
			panic("cache/lru/node.rebuild() tail node cannot have .next == nil")
		}
		c.tail = node.next
		node.next.previous = nil

		node.previous = c.head
		node.next = nil

		c.head.next = node
		c.head = node
	default:
		if node.next == nil || node.previous == nil {
			panic("cache/lru/node.rebuild() node cannot have .next or .previous == nil")
		}

		node.previous.next = node.next
		node.next.previous = node.previous

		node.previous = c.head
		node.next = nil

		c.head.next = node
		c.head = node
	}
}

func (c *cache) dropTail() {
	var tail = c.tail
	c.tail = tail.next

	delete(c.indexes, *tail.key)
	tail.drop()
	c.len--
}

func (c *cache) Get(key string) ([][]byte, bool) {
	if key == "" {
		return nil, false
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	var node, ok = c.indexes[key]
	if !ok {
		return nil, false
	}

	c.rebuild(node)

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

	var _node, ok = c.indexes[key]
	switch {
	case ok && !c.mutable:
		return erres.AlreadyExist.Extend(0).SetDescription("given key already exist")
	case ok && c.mutable:
		_node.value = value
		c.rebuild(_node)
		return nil
	case !ok:
		var node = new(node)
		node.key = &key
		node.value = value
		c.indexes[key] = node

		if c.len == c.size {
			c.dropTail()
		}

		c.build(node)
	}

	return nil
}

func NewCache(size uint, mutable bool, acceptNilValue bool) (Cache, error) {
	if size == 0 {
		return nil, erres.InvalidArgument.Extend(0).SetDescription("size cannot be zero")
	}

	var cache = new(cache)

	cache.mutex = new(sync.Mutex)
	cache.indexes = make(map[string]*node, size)

	cache.size = size
	cache.mutable = mutable
	cache.acceptNilValue = acceptNilValue

	return cache, nil
}
