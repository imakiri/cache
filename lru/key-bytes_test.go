package lru

import (
	"testing"
)

func setKB(cache KeyBytes, i int) error {
	return cache.Set(keys[i], values[i])
}

func getKB(cache KeyBytes, i int) ([][]byte, bool) {
	return cache.Get(keys[i])
}

func TestKB1f(t *testing.T) {
	var cache, err = NewKeyBytesCache(1, false)
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

	if err = setKB(cache, 0); err != nil {
		t.Error(err)
	}
	if value, ok = getKB(cache, 0); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 0) {
		t.Error("not equal")
	}

	if err = setKB(cache, 1); err != nil {
		t.Error(err)
	}
	if err = setKB(cache, 2); err != nil {
		t.Error(err)
	}
	if value, ok = getKB(cache, 2); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 2) {
		t.Error("not equal")
	}

	if err = setKB(cache, 3); err != nil {
		t.Error(err)
	}
	if err = setKB(cache, 4); err != nil {
		t.Error(err)
	}
	if err = setKB(cache, 5); err != nil {
		t.Error(err)
	}
	if value, ok = getKB(cache, 5); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 5) {
		t.Error("not equal")
	}

	if err = setKB(cache, 6); err != nil {
		t.Error(err)
	}
	if err = setKB(cache, 7); err != nil {
		t.Error(err)
	}
	if err = setKB(cache, 8); err != nil {
		t.Error(err)
	}
	if err = setKB(cache, 9); err != nil {
		t.Error(err)
	}
	if value, ok = getKB(cache, 9); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 9) {
		t.Error("not equal")
	}

}

func TestKB3f(t *testing.T) {
	var cache, err = NewKeyBytesCache(3, false)
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

	if err = setKB(cache, 0); err != nil {
		t.Error(err)
	}

	if value, ok = getKB(cache, 0); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 0) {
		t.Error("not equal")
	}

	if err = setKB(cache, 1); err != nil {
		t.Error(err)
	}
	if err = setKB(cache, 2); err != nil {
		t.Error(err)
	}

	if value, ok = getKB(cache, 0); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 0) {
		t.Error("not equal")
	}
	if value, ok = getKB(cache, 1); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 1) {
		t.Error("not equal")
	}
	if value, ok = getKB(cache, 2); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 2) {
		t.Error("not equal")
	}

	if err = setKB(cache, 3); err != nil {
		t.Error(err)
	}

	if value, ok = getKB(cache, 0); ok {
		t.Error("retrieved non-existent value")
	}

	if err = setKB(cache, 4); err != nil {
		t.Error(err)
	}
	if err = setKB(cache, 5); err != nil {
		t.Error(err)
	}

	if value, ok = getKB(cache, 5); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 5) {
		t.Error("not equal")
	}

	if err = setKB(cache, 6); err != nil {
		t.Error(err)
	}
	if err = setKB(cache, 7); err != nil {
		t.Error(err)
	}
	if err = setKB(cache, 8); err != nil {
		t.Error(err)
	}
	if err = setKB(cache, 9); err != nil {
		t.Error(err)
	}

	if value, ok = getKB(cache, 9); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 9) {
		t.Error("not equal")
	}
	if value, ok = getKB(cache, 8); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 8) {
		t.Error("not equal")
	}
	if value, ok = getKB(cache, 7); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 7) {
		t.Error("not equal")
	}
}
