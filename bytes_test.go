package immutable

import (
	"strconv"
	"testing"
)

func assertBytesEqual(t *testing.T, want, got Bytes) {
	if !got.Equal(want) {
		t.Errorf("assert bytes equal failed: got %s, want %s", got.p, want.p)
	}
}

func assertIntEqual(t *testing.T, want, got int) {
	if got != want {
		t.Errorf("assert int equal failed: got %d, want %d", got, want)
	}
}

func assertTrue(t *testing.T, got bool) {
	if got != true {
		t.Errorf("assert bool equal failed: got %v, want true", got)
	}
}

func assertFalse(t *testing.T, got bool) {
	if got != false {
		t.Errorf("assert bool equal failed: got %v, want false", got)
	}
}

var (
	nilBytes   = NewBytes(nil, false)
	emptyBytes = NewBytes(make([]byte, 0), false)
)

func Test_JoinBytes(t *testing.T) {
	assertBytesEqual(t, nilBytes, JoinBytes(BytesFromString(",")))
	assertBytesEqual(t, nilBytes, JoinBytes(BytesFromString(","), emptyBytes))
	assertBytesEqual(t, BytesFromString("1"), JoinBytes(BytesFromString(","), BytesFromString("1")))
	assertBytesEqual(
		t,
		BytesFromString("1,2,3"),
		JoinBytes(
			BytesFromString(","),
			BytesFromString("1"),
			BytesFromString("2"),
			BytesFromString("3"),
		),
	)
}

func Test_JoinBytesR(t *testing.T) {
	assertBytesEqual(t, nilBytes, JoinBytesR([]byte(",")))
	assertBytesEqual(t, nilBytes, JoinBytesR([]byte(","), emptyBytes))
	assertBytesEqual(t, BytesFromString("1"), JoinBytesR([]byte(","), BytesFromString("1")))
	assertBytesEqual(
		t,
		BytesFromString("1,2,3"),
		JoinBytesR(
			[]byte(","),
			BytesFromString("1"),
			BytesFromString("2"),
			BytesFromString("3"),
		),
	)
}

func TestBytes_Append(t *testing.T) {
	a := NewBytes([]byte{'1', '2', '3'}, false)
	b := NewBytes([]byte{'4', '5', '6'}, false)
	c := NewBytes([]byte{'7', '8', '9'}, false)
	d := a.Append(b, c)
	assertBytesEqual(t, BytesFromString("123456789"), d)
	e := a.Append()
	assertBytesEqual(t, BytesFromString("123"), e)
}

func TestBytes_Compare(t *testing.T) {
	ar := []byte{'1', '2', '3'}
	br := []byte{'4', '5', '6'}
	cr := []byte{'7', '8', '9'}
	a := NewBytes(ar, false)
	b := NewBytes(br, false)
	c := NewBytes(cr, false)

	assertIntEqual(t, 1, b.Compare(a))
	assertIntEqual(t, 0, b.Compare(b))
	assertIntEqual(t, -1, b.Compare(c))
}

func TestBytes_CompareR(t *testing.T) {
	ar := []byte{'1', '2', '3'}
	br := []byte{'4', '5', '6'}
	cr := []byte{'7', '8', '9'}
	b := NewBytes(br, false)

	assertIntEqual(t, 1, b.CompareR(ar))
	assertIntEqual(t, 0, b.CompareR(br))
	assertIntEqual(t, -1, b.CompareR(cr))
}

func TestBytes_Contains(t *testing.T) {
	ar := []byte{'1', '2', '3'}
	br := []byte{'4', '5', '6'}
	a := NewBytes(ar, false)
	b := NewBytes(br, false)

	assertTrue(t, a.Contains(a))
	assertTrue(t, a.Contains(nilBytes))
	assertFalse(t, a.Contains(b))
	assertFalse(t, nilBytes.Contains(b))
}

func TestBytes_ContainsAny(t *testing.T) {
	// todo
}

func TestBytes_ContainsR(t *testing.T) {
	ar := []byte{'1', '2', '3'}
	br := []byte{'4', '5', '6'}
	a := NewBytes(ar, false)

	assertTrue(t, a.ContainsR(ar))
	assertTrue(t, a.ContainsR(nil))
	assertFalse(t, a.ContainsR(br))
	assertFalse(t, nilBytes.ContainsR(br))
}

func TestBytes_ContainsRune(t *testing.T) {
	// todo
}

func TestBytes_Copy(t *testing.T) {
	shortBuf := make([]byte, 2)
	assertIntEqual(t, 2, BytesFromString("123").Copy(shortBuf))
	buf := make([]byte, 3)
	assertIntEqual(t, 3, BytesFromString("123").Copy(buf))
	longBuf := make([]byte, 4)
	assertIntEqual(t, 3, BytesFromString("123").Copy(longBuf))
}

func TestBytes_Count(t *testing.T) {
	assertIntEqual(t, 0, nilBytes.Count(BytesFromString("0")))
	assertIntEqual(t, 0, BytesFromString("123").Count(BytesFromString("0")))
	assertIntEqual(t, 1, BytesFromString("123").Count(BytesFromString("1")))
	assertIntEqual(t, 2, BytesFromString("1231").Count(BytesFromString("1")))
	assertIntEqual(t, 4, BytesFromString("123").Count(nilBytes))
}

func TestBytes_CountR(t *testing.T) {
	assertIntEqual(t, 0, nilBytes.CountR([]byte{'0'}))
	assertIntEqual(t, 0, BytesFromString("123").CountR([]byte{'0'}))
	assertIntEqual(t, 1, BytesFromString("123").CountR([]byte{'1'}))
	assertIntEqual(t, 2, BytesFromString("1231").CountR([]byte{'1'}))
	assertIntEqual(t, 4, BytesFromString("123").CountR(nil))
}

func TestBytes_Equal(t *testing.T) {
	assertTrue(t, nilBytes.Equal(emptyBytes))
	assertTrue(t, BytesFromString("123").Equal(BytesFromString("123")))
	assertFalse(t, BytesFromString("123").Equal(BytesFromString("456")))
}

func TestBytes_EqualFold(t *testing.T) {
	assertTrue(t, nilBytes.EqualFold(emptyBytes))
	assertTrue(t, BytesFromString("abc").EqualFold(BytesFromString("ABC")))
	assertFalse(t, BytesFromString("123").EqualFold(BytesFromString("abc")))
}

func TestBytes_EqualFoldR(t *testing.T) {
	assertTrue(t, nilBytes.EqualFoldR(nil))
	assertTrue(t, BytesFromString("abc").EqualFoldR([]byte("ABC")))
	assertFalse(t, BytesFromString("123").EqualFoldR([]byte("abc")))
}

func TestBytes_EqualR(t *testing.T) {
	assertTrue(t, nilBytes.EqualR(nil))
	assertTrue(t, BytesFromString("123").EqualR([]byte("123")))
	assertFalse(t, BytesFromString("123").EqualR([]byte("456")))
}

func TestBytes_Fields(t *testing.T) {
	// todo
}

func TestBytes_FieldsFunc(t *testing.T) {
	// todo
}

func TestBytes_HasPrefix(t *testing.T) {
	// todo
}

func TestBytes_HasPrefixR(t *testing.T) {
	// todo
}

func TestBytes_HasSuffix(t *testing.T) {
	// todo
}

func TestBytes_HasSuffixR(t *testing.T) {
	// todo
}

func TestBytes_Index(t *testing.T) {
	// todo
}

func TestBytes_IndexAny(t *testing.T) {
	// todo
}

func TestBytes_IndexByte(t *testing.T) {
	// todo
}

func TestBytes_IndexFunc(t *testing.T) {
	// todo
}

func TestBytes_IndexR(t *testing.T) {
	// todo
}

func TestBytes_IndexRune(t *testing.T) {
	// todo
}

func TestBytes_LastIndex(t *testing.T) {
	// todo
}

func TestBytes_LastIndexAny(t *testing.T) {
	// todo
}

func TestBytes_LastIndexByte(t *testing.T) {
	// todo
}

func TestBytes_LastIndexFunc(t *testing.T) {
	// todo
}

func TestBytes_LastIndexR(t *testing.T) {
	// todo
}

func TestBytes_Len(t *testing.T) {
	// todo
}

func TestBytes_Map(t *testing.T) {
	// todo
}

func TestBytes_MarshalBinary(t *testing.T) {
	cases := [][]byte{
		nil,
		make([]byte, 0, 1),
		[]byte("message"),
	}
	for idx, c := range cases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			b1 := NewBytes(c, false)
			data, err := b1.MarshalBinary()
			if err != nil {
				t.Fatal(err)
			}
			b2 := Bytes{}
			if err := b2.UnmarshalBinary(data); err != nil {
				t.Fatal(err)
			}
			assertBytesEqual(t, b1, b2)
		})
	}
}

func TestBytes_MarshalText(t *testing.T) {
	cases := [][]byte{
		nil,
		make([]byte, 0, 1),
		[]byte("message"),
	}
	for idx, c := range cases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			b1 := NewBytes(c, false)
			data, err := b1.MarshalText()
			if err != nil {
				t.Fatal(err)
			}
			b2 := Bytes{}
			if err := b2.UnmarshalText(data); err != nil {
				t.Fatal(err)
			}
			assertBytesEqual(t, b1, b2)
		})
	}
}

func TestBytes_Reader(t *testing.T) {
	// todo
}

func TestBytes_Repeat(t *testing.T) {
	// todo
}

func TestBytes_Replace(t *testing.T) {
	// todo
}

func TestBytes_ReplaceAll(t *testing.T) {
	// todo
}

func TestBytes_ReplaceAllR(t *testing.T) {
	// todo
}

func TestBytes_ReplaceR(t *testing.T) {
	// todo
}

func TestBytes_Runes(t *testing.T) {
	// todo
}

func TestBytes_Slice(t *testing.T) {
	// todo
}

func TestBytes_Split(t *testing.T) {
	// todo
}

func TestBytes_SplitAfter(t *testing.T) {
	// todo
}

func TestBytes_SplitAfterN(t *testing.T) {
	// todo
}

func TestBytes_SplitAfterNR(t *testing.T) {
	// todo
}

func TestBytes_SplitAfterR(t *testing.T) {
	// todo
}

func TestBytes_SplitN(t *testing.T) {
	// todo
}

func TestBytes_SplitNR(t *testing.T) {
	// todo
}

func TestBytes_SplitR(t *testing.T) {
	// todo
}

func TestBytes_Title(t *testing.T) {
	// todo
}

func TestBytes_ToLower(t *testing.T) {
	// todo
}

func TestBytes_ToLowerSpecial(t *testing.T) {
	// todo
}

func TestBytes_ToTitle(t *testing.T) {
	// todo
}

func TestBytes_ToTitleSpecial(t *testing.T) {
	// todo
}

func TestBytes_ToUpper(t *testing.T) {
	// todo
}

func TestBytes_ToUpperSpecial(t *testing.T) {
	// todo
}

func TestBytes_ToValidUTF8(t *testing.T) {
	// todo
}

func TestBytes_ToValidUTF8R(t *testing.T) {
	// todo
}

func TestBytes_Trim(t *testing.T) {
	// todo
}

func TestBytes_TrimFunc(t *testing.T) {
	// todo
}

func TestBytes_TrimLeft(t *testing.T) {
	// todo
}

func TestBytes_TrimLeftFunc(t *testing.T) {
	// todo
}

func TestBytes_TrimPrefix(t *testing.T) {
	// todo
}

func TestBytes_TrimPrefixR(t *testing.T) {
	// todo
}

func TestBytes_TrimRight(t *testing.T) {
	// todo
}

func TestBytes_TrimRightFunc(t *testing.T) {
	// todo
}

func TestBytes_TrimSpace(t *testing.T) {
	// todo
}

func TestBytes_TrimSuffix(t *testing.T) {
	// todo
}

func TestBytes_TrimSuffixR(t *testing.T) {
	// todo
}

func TestNewBytes(t *testing.T) {
	// todo
}

func TestNewBytesSlice(t *testing.T) {
	// todo
}
