package lru

import (
	"github.com/imakiri/erres"
	"sync"
)

type general struct {
	mutex *sync.Mutex

	indexes map[string]Node
	head    Node
	tail    Node

	len  uint
	size uint
}

func (c *general) lead(node Node) {
	if node.GetPrevious() != nil || node.GetNext() != nil {
		panic("cannot assign node with links as head")
	}

	c.head.SetNext(node)
	node.SetPrevious(c.head)
	c.head = node
}

func (c *general) drop(node Node) {
	switch {
	case node == c.head && node == c.tail:
		break
	case node == c.tail:
		c.tail = node.GetNext()
		node.GetNext().SetPrevious(nil)
		node.SetNext(nil)
	case node.GetPrevious() != nil && node.GetNext() != nil:
		node.GetPrevious().SetNext(node.GetNext())
		node.GetNext().SetPrevious(node.GetPrevious())
		node.SetNext(nil)
		node.SetPrevious(nil)
	default:
		panic("drop default")
	}
}

func (c *general) build(node Node) {
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

func (c *general) rebuild(node Node) {
	switch {
	case node == nil:
		panic("node cannot be nil")
	case node == c.head:
		return
	}

	c.drop(node)
	c.lead(node)
}

func (c *general) delete() {
	var tail = c.tail
	c.drop(tail)
	delete(c.indexes, tail.GetKey())
	tail.Drop()
	c.len--
}

func (c *general) Get(key string) (Node, bool) {
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
	return node, true
}

func (c *general) Set(node Node) error {
	switch {
	case node.GetKey() == "":
		return erres.InvalidArgument.Extend(0).SetDescription("key string cannot be empty")
	case node == nil:
		return erres.NilArgument.Extend(0).SetDescription("value cannot be nil")
	}

	var key = node.GetKey()

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.indexes[key] = node
	if c.len == c.size {
		c.delete()
	}
	c.build(node)

	return nil
}
