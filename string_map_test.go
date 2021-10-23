package immutable

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestNewStringMap(t *testing.T) {
	// todo
}

func TestStringMap_Copy(t *testing.T) {
	// todo
}

func TestStringMap_Default(t *testing.T) {
	// todo
}

func TestStringMap_Get(t *testing.T) {
	// todo
}

func TestStringMap_Has(t *testing.T) {
	// todo
}

func TestStringMap_Keys(t *testing.T) {
	// todo
}

func TestStringMap_Len(t *testing.T) {
	// todo
}

func TestStringMap_MarshalBinary(t *testing.T) {
	cases := []map[string]string{
		nil,
		make(map[string]string, 10),
		{"a": "1"},
		{"a": "1", "b": "2"},
		{"a": strings.Repeat("a", math.MaxUint16+1)},
	}
	for idx, c := range cases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			sm1 := NewStringMap(c, false)
			data, err := sm1.MarshalBinary()
			if err != nil {
				t.Fatal(err)
			}
			sm2 := StringMap{}
			if err := sm2.UnmarshalBinary(data); err != nil {
				t.Fatal(err)
			}
			if !sm1.IsEqual(sm2) {
				t.Fatal("values are not equal")
			}
		})
	}
}

func TestStringMap_MarshalText(t *testing.T) {
	cases := []map[string]string{
		nil,
		make(map[string]string, 10),
		{"a": "1"},
		{"a": "1", "b": "2"},
	}
	for idx, c := range cases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			sm1 := NewStringMap(c, false)
			data, err := sm1.MarshalText()
			if err != nil {
				t.Fatal(err)
			}
			sm2 := StringMap{}
			if err := sm2.UnmarshalText(data); err != nil {
				t.Fatal(err)
			}
			if !sm1.IsEqual(sm2) {
				t.Fatal("values are not equal")
			}
		})
	}
}

func TestStringMap_Merge(t *testing.T) {
	// todo
}

func TestStringMap_With(t *testing.T) {
	// todo
}

func TestStringMap_Without(t *testing.T) {
	// todo
}

func TestStringMap_merge(t *testing.T) {
	// todo
}
