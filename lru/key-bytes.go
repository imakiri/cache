package lru

import (
	"github.com/imakiri/erres"
	"sync"
)

type keyBytes struct {
	mutex *sync.Mutex

	indexes map[string]*node
	head    *node
	tail    *node

	len            uint
	size           uint
	acceptNilValue bool
}

func (c *keyBytes) lead(node *node) {
	if node.previous != nil || node.next != nil {
		panic("cannot assign node with links as head")
	}

	c.head.next = node
	node.previous = c.head
	c.head = node
}

func (c *keyBytes) drop(node *node) {
	switch {
	case node == c.head && node == c.tail:
		break
	case node == c.tail:
		c.tail = node.next
		node.next.previous = nil
		node.next = nil
	case node.previous != nil && node.next != nil:
		node.previous.next = node.next
		node.next.previous = node.previous
		node.next = nil
		node.previous = nil
	default:
		panic("drop default")
	}
}

func (c *keyBytes) build(node *node) {
	if node == nil {
		panic("node cannot be nil")
	}

	switch {
	case c.len == 0:
		c.tail = node
		c.head = node
	default:
		c.lead(node)
	}

	c.len++
}

func (c *keyBytes) rebuild(node *node) {
	switch {
	case node == nil:
		panic("node cannot be nil")
	case node == c.head:
		return
	}

	c.drop(node)
	c.lead(node)
}

func (c *keyBytes) delete() {
	var tail = c.tail
	c.drop(tail)
	delete(c.indexes, *tail.key)
	tail.key = nil
	tail.value = nil
	c.len--
}

func (c *keyBytes) Get(key string) ([][]byte, bool) {
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

func (c *keyBytes) Set(key string, value [][]byte) error {
	switch {
	case key == "":
		return erres.InvalidArgument.Extend(0).SetDescription("key string cannot be empty")
	case value == nil && !c.acceptNilValue:
		return erres.NilArgument.Extend(0).SetDescription("value cannot be nil")
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	var _node, ok = c.indexes[key]
	if ok {
		_node.value = value
		c.rebuild(_node)
	}

	var node = new(node)
	node.key = &key
	node.value = value

	c.indexes[key] = node
	if c.len == c.size {
		c.delete()
	}
	c.build(node)

	return nil
}
