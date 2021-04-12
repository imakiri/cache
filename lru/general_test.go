package lru

import (
	"testing"
)

func setG(cache General, i int) error {
	return cache.Set(&generalNode{
		next:     nil,
		previous: nil,
		key:      &keys[i],
		value:    values[i],
	})
}

func getG(cache General, i int) ([][]byte, bool) {
	var ok bool
	var n Node
	var _n *generalNode

	if n, ok = cache.Get(keys[i]); !ok {
		return nil, false
	}
	if _n, ok = n.(*generalNode); !ok {
		return nil, false
	}

	return _n.value, true
}

func TestG1(t *testing.T) {
	var cache, err = NewGeneralCache(1)
	if err != nil {
		t.Error(err)
	}

	for _, k := range keys {
		var _, ok = cache.Get(k)
		if ok {
			t.Error("retrieved non-existent value")
		}
	}

	var value [][]byte
	var ok bool

	if err = setG(cache, 0); err != nil {
		t.Error(err)
	}
	if value, ok = getG(cache, 0); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 0) {
		t.Error("not equal")
	}

	if err = setG(cache, 1); err != nil {
		t.Error(err)
	}
	if err = setG(cache, 2); err != nil {
		t.Error(err)
	}
	if value, ok = getG(cache, 2); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 2) {
		t.Error("not equal")
	}

	if err = setG(cache, 3); err != nil {
		t.Error(err)
	}
	if err = setG(cache, 4); err != nil {
		t.Error(err)
	}
	if err = setG(cache, 5); err != nil {
		t.Error(err)
	}
	if value, ok = getG(cache, 5); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 5) {
		t.Error("not equal")
	}

	if err = setG(cache, 6); err != nil {
		t.Error(err)
	}
	if err = setG(cache, 7); err != nil {
		t.Error(err)
	}
	if err = setG(cache, 8); err != nil {
		t.Error(err)
	}
	if err = setG(cache, 9); err != nil {
		t.Error(err)
	}
	if value, ok = getG(cache, 9); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 9) {
		t.Error("not equal")
	}

}

func TestG3(t *testing.T) {
	var cache, err = NewGeneralCache(3)
	if err != nil {
		t.Error(err)
	}

	// Getting non-existent values
	for _, k := range keys {
		var _, ok = cache.Get(k)
		if ok {
			t.Error("retrieved non-existent value")
		}
	}

	var value [][]byte
	var ok bool

	if err = setG(cache, 0); err != nil {
		t.Error(err)
	}

	if value, ok = getG(cache, 0); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 0) {
		t.Error("not equal")
	}

	if err = setG(cache, 1); err != nil {
		t.Error(err)
	}
	if err = setG(cache, 2); err != nil {
		t.Error(err)
	}

	if value, ok = getG(cache, 0); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 0) {
		t.Error("not equal")
	}
	if value, ok = getG(cache, 1); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 1) {
		t.Error("not equal")
	}
	if value, ok = getG(cache, 2); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 2) {
		t.Error("not equal")
	}

	if err = setG(cache, 3); err != nil {
		t.Error(err)
	}

	if value, ok = getG(cache, 0); ok {
		t.Error("retrieved non-existent value")
	}

	if err = setG(cache, 4); err != nil {
		t.Error(err)
	}
	if err = setG(cache, 5); err != nil {
		t.Error(err)
	}

	if value, ok = getG(cache, 5); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 5) {
		t.Error("not equal")
	}

	if err = setG(cache, 6); err != nil {
		t.Error(err)
	}
	if err = setG(cache, 7); err != nil {
		t.Error(err)
	}
	if err = setG(cache, 8); err != nil {
		t.Error(err)
	}
	if err = setG(cache, 9); err != nil {
		t.Error(err)
	}

	if value, ok = getG(cache, 9); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 9) {
		t.Error("not equal")
	}
	if value, ok = getG(cache, 8); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 8) {
		t.Error("not equal")
	}
	if value, ok = getG(cache, 7); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 7) {
		t.Error("not equal")
	}
}
