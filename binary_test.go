package immutable

import (
	"strconv"
	"testing"
)

func TestVarint(t *testing.T) {
	type testCase struct {
		val uint64
		res []byte
	}
	cases := []testCase{
		{0, []byte{0}},
		{1, []byte{1}},
		{0xFC, []byte{0xFC}},
		{0xFD, []byte{0xFD, 0xFD, 0x00}},
		{0xFE, []byte{0xFD, 0xFE, 0x00}},
		{0xFF, []byte{0xFD, 0xFF, 0x00}},
		{0xFFFF, []byte{0xFD, 0xFF, 0xFF}},
		{0x00010000, []byte{0xFE, 0x00, 0x00, 0x01, 0x00}},
		{0xFFFFFFFF, []byte{0xFE, 0xFF, 0xFF, 0xFF, 0xFF}},
		{0x0000000100000000, []byte{0xFF, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00}},
		{0xFFFFFFFFFFFFFFFF, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
	}
	for _, c := range cases {
		t.Run(strconv.FormatUint(c.val, 16), func(t *testing.T) {
			size := varintSize(c.val)
			if size != len(c.res) {
				t.Fatal("invalid varint size")
			}
			buf := make([]byte, size)
			if n := putVarint(buf, c.val); n != size {
				t.Fatal("invalid put varint size")
			}
			if string(buf) != string(c.res) {
				t.Fatal("invalid put varint result")
			}
			val, n := getVarint(buf)
			if n != size {
				t.Fatal("invalid get varint size")
			}
			if val != c.val {
				t.Fatal("invalid get varint result")
			}
		})
	}
}
