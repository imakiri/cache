package lru

type Node interface {
	GetNext() Node
	SetNext(Node)
	GetPrevious() Node
	SetPrevious(Node)
	GetKey() string
	Drop()
}

type node struct {
	next     *node
	previous *node
	key      *string
	value    [][]byte
}

type generalNode struct {
	next     *generalNode
	previous *generalNode
	key      *string
	value    [][]byte
}

func (n *generalNode) GetNext() Node {
	if n.next == nil {
		return nil
	}
	return n.next
}

func (n *generalNode) SetNext(no Node) {
	var node, ok = no.(*generalNode)
	if ok {
		n.next = node
	} else {
		n.next = nil
	}
}

func (n *generalNode) GetPrevious() Node {
	if n.previous == nil {
		return nil
	}
	return n.previous
}

func (n *generalNode) SetPrevious(no Node) {
	var node, ok = no.(*generalNode)
	if ok {
		n.previous = node
	} else {
		n.previous = nil
	}
}

func (n *generalNode) GetKey() string {
	return *n.key
}

func (n *generalNode) Drop() {
	n.key = nil
	n.value = nil
}
