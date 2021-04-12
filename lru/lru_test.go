package lru

import "bytes"

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

func equal(values [][][]byte, value [][]byte, i int) (b bool) {
	defer func() {
		if recover() != nil {
			b = false
		}
	}()

	return bytes.Equal(values[i][0], value[0])
}
