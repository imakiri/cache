package lru

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var cache, err = NewCache(1, false, false)
	if err != nil {
		t.Error(err)
	}

	var keys = []string{
		"1 asd",
		"2 ggd",
		"3 erd",
		"4 reg",
		"5 erf",
		"6 ard",
		"7 itd",
	}

	var values = [][][]byte{
		{[]byte("1")},
		{[]byte("2")},
		{[]byte("3")},
		{[]byte("4")},
		{[]byte("5")},
		{[]byte("6")},
		{[]byte("7")},
	}

	for i, k := range keys {
		if err = cache.Set(k, values[i]); err != nil {
			t.Error(err)
		}
	}

	for _, k := range keys {
		var vls, ok = cache.Get(k)
		if !ok && k == keys[6] {
			t.Error("not ok")
		}

		fmt.Println(vls)
	}
}
