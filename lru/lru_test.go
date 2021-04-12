package lru

import (
	"bytes"
	"testing"
)

var keys = []string{
	"1 asd",
	"2 ggd",
	"3 erd",
	"4 reg",
	"5 erf",
	"6 ard",
	"7 itd",
	"8 yut",
	"9 qkf",
	"10 yb",
	"11 pz",
}

var values = [][][]byte{
	{[]byte("1 erf")},
	{[]byte("2 gsg")},
	{[]byte("3 pma")},
	{[]byte("4 ert")},
	{[]byte("5 mze")},
	{[]byte("6 qzu")},
	{[]byte("7 ase")},
	{[]byte("8 jjh")},
	{[]byte("9 haq")},
	{[]byte("10 qq")},
	{[]byte("11 ak")},
}

func set(cache Cache, i int) error {
	return cache.Set(keys[i], values[i])
}

func get(cache Cache, i int) ([][]byte, bool) {
	return cache.Get(keys[i])
}

func equal(values [][][]byte, value [][]byte, i int) (b bool) {
	defer func() {
		if recover() != nil {
			b = false
		}
	}()

	return bytes.Equal(values[i][0], value[0])
}

func Test1ff(t *testing.T) {
	var cache, err = NewCache(1, false, false)
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

	if err = set(cache, 0); err != nil {
		t.Error(err)
	}
	if value, ok = get(cache, 0); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 0) {
		t.Error("not equal")
	}

	if err = set(cache, 1); err != nil {
		t.Error(err)
	}
	if err = set(cache, 2); err != nil {
		t.Error(err)
	}
	if value, ok = get(cache, 2); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 2) {
		t.Error("not equal")
	}

	if err = set(cache, 3); err != nil {
		t.Error(err)
	}
	if err = set(cache, 4); err != nil {
		t.Error(err)
	}
	if err = set(cache, 5); err != nil {
		t.Error(err)
	}
	if value, ok = get(cache, 5); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 5) {
		t.Error("not equal")
	}

	if err = set(cache, 6); err != nil {
		t.Error(err)
	}
	if err = set(cache, 7); err != nil {
		t.Error(err)
	}
	if err = set(cache, 8); err != nil {
		t.Error(err)
	}
	if err = set(cache, 9); err != nil {
		t.Error(err)
	}
	if value, ok = get(cache, 9); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 9) {
		t.Error("not equal")
	}

}

func Test3ff(t *testing.T) {
	var cache, err = NewCache(3, false, false)
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

	if err = set(cache, 0); err != nil {
		t.Error(err)
	}

	if value, ok = get(cache, 0); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 0) {
		t.Error("not equal")
	}

	if err = set(cache, 1); err != nil {
		t.Error(err)
	}
	if err = set(cache, 2); err != nil {
		t.Error(err)
	}

	if value, ok = get(cache, 0); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 0) {
		t.Error("not equal")
	}
	if value, ok = get(cache, 1); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 1) {
		t.Error("not equal")
	}
	if value, ok = get(cache, 2); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 2) {
		t.Error("not equal")
	}

	if err = set(cache, 3); err != nil {
		t.Error(err)
	}

	if value, ok = get(cache, 0); ok {
		t.Error("retrieved non-existent value")
	}

	if err = set(cache, 4); err != nil {
		t.Error(err)
	}
	if err = set(cache, 5); err != nil {
		t.Error(err)
	}

	if value, ok = get(cache, 5); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 5) {
		t.Error("not equal")
	}

	if err = set(cache, 6); err != nil {
		t.Error(err)
	}
	if err = set(cache, 7); err != nil {
		t.Error(err)
	}
	if err = set(cache, 8); err != nil {
		t.Error(err)
	}
	if err = set(cache, 9); err != nil {
		t.Error(err)
	}

	if value, ok = get(cache, 9); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 9) {
		t.Error("not equal")
	}
	if value, ok = get(cache, 8); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 8) {
		t.Error("not equal")
	}
	if value, ok = get(cache, 7); !ok {
		t.Error("not ok")
	}
	if !equal(values, value, 7) {
		t.Error("not equal")
	}
}
