package lru

import (
	"github.com/imakiri/erres"
	"sync"
)

type KeyBytes interface {
	Get(key string) ([][]byte, bool)
	Set(key string, value [][]byte) error
}

type General interface {
	Get(key string) (Node, bool)
	Set(node Node) error
}

func NewKeyBytesCache(size uint, acceptNilValue bool) (KeyBytes, error) {
	if size == 0 {
		return nil, erres.InvalidArgument.Extend(0).SetDescription("size cannot be zero")
	}

	var cache = new(keyBytes)

	cache.mutex = new(sync.Mutex)
	cache.indexes = make(map[string]*node, size)

	cache.size = size
	cache.acceptNilValue = acceptNilValue

	return cache, nil
}

func NewGeneralCache(size uint) (General, error) {
	if size == 0 {
		return nil, erres.InvalidArgument.Extend(0).SetDescription("size cannot be zero")
	}

	var cache = new(general)

	cache.mutex = new(sync.Mutex)
	cache.indexes = make(map[string]Node, size)

	cache.size = size

	return cache, nil
}
