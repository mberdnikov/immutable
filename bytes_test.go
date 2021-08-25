package immutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func assertBytes(got, want Bytes, t *testing.T) {
	if !got.Equal(want) {
		t.Errorf("got %s, want %s", got.p, want.p)
	}
}

var (
	nilBytes   = NewBytes(nil, false)
	emptyBytes = NewBytes(make([]byte, 0), false)
)

func Test_JoinBytes(t *testing.T) {
	assertBytes(JoinBytes(BytesFromString(",")), nilBytes, t)
	assertBytes(JoinBytes(BytesFromString(","), emptyBytes), nilBytes, t)
	assertBytes(JoinBytes(BytesFromString(","), BytesFromString("1")), BytesFromString("1"), t)
	assertBytes(
		JoinBytes(
			BytesFromString(","),
			BytesFromString("1"),
			BytesFromString("2"),
			BytesFromString("3"),
		),
		BytesFromString("1,2,3"),
		t,
	)
}

func Test_JoinBytesR(t *testing.T) {
	assertBytes(JoinBytesR([]byte(",")), nilBytes, t)
	assertBytes(JoinBytesR([]byte(","), emptyBytes), nilBytes, t)
	assertBytes(JoinBytesR([]byte(","), BytesFromString("1")), BytesFromString("1"), t)
	assertBytes(
		JoinBytesR(
			[]byte(","),
			BytesFromString("1"),
			BytesFromString("2"),
			BytesFromString("3"),
		),
		BytesFromString("1,2,3"),
		t,
	)
}

func TestBytes_Append(t *testing.T) {
	a := NewBytes([]byte{'1', '2', '3'}, false)
	b := NewBytes([]byte{'4', '5', '6'}, false)
	c := NewBytes([]byte{'7', '8', '9'}, false)
	d := a.Append(b, c)
	assertBytes(d, BytesFromString("123456789"), t)
	e := a.Append()
	assertBytes(e, BytesFromString("123"), t)
}

func TestBytes_Compare(t *testing.T) {
	ar := []byte{'1', '2', '3'}
	br := []byte{'4', '5', '6'}
	cr := []byte{'7', '8', '9'}
	a := NewBytes(ar, false)
	b := NewBytes(br, false)
	c := NewBytes(cr, false)

	assert.Equal(t, 1, b.Compare(a))
	assert.Equal(t, 0, b.Compare(b))
	assert.Equal(t, -1, b.Compare(c))
}

func TestBytes_CompareR(t *testing.T) {
	ar := []byte{'1', '2', '3'}
	br := []byte{'4', '5', '6'}
	cr := []byte{'7', '8', '9'}
	b := NewBytes(br, false)

	assert.Equal(t, 1, b.CompareR(ar))
	assert.Equal(t, 0, b.CompareR(br))
	assert.Equal(t, -1, b.CompareR(cr))
}

func TestBytes_Contains(t *testing.T) {
	ar := []byte{'1', '2', '3'}
	br := []byte{'4', '5', '6'}
	a := NewBytes(ar, false)
	b := NewBytes(br, false)

	assert.True(t, a.Contains(a))
	assert.True(t, a.Contains(nilBytes))
	assert.False(t, a.Contains(b))
	assert.False(t, nilBytes.Contains(b))
}

func TestBytes_ContainsAny(t *testing.T) {
	// todo
}

func TestBytes_ContainsR(t *testing.T) {
	ar := []byte{'1', '2', '3'}
	br := []byte{'4', '5', '6'}
	a := NewBytes(ar, false)

	assert.True(t, a.ContainsR(ar))
	assert.True(t, a.ContainsR(nil))
	assert.False(t, a.ContainsR(br))
	assert.False(t, nilBytes.ContainsR(br))
}

func TestBytes_ContainsRune(t *testing.T) {
	// todo
}

func TestBytes_Copy(t *testing.T) {
	shortBuf := make([]byte, 2)
	assert.Equal(t, 2, BytesFromString("123").Copy(shortBuf))
	buf := make([]byte, 3)
	assert.Equal(t, 3, BytesFromString("123").Copy(buf))
	longBuf := make([]byte, 4)
	assert.Equal(t, 3, BytesFromString("123").Copy(longBuf))
}

func TestBytes_Count(t *testing.T) {
	assert.Equal(t, 0, nilBytes.Count(BytesFromString("0")))
	assert.Equal(t, 0, BytesFromString("123").Count(BytesFromString("0")))
	assert.Equal(t, 1, BytesFromString("123").Count(BytesFromString("1")))
	assert.Equal(t, 2, BytesFromString("1231").Count(BytesFromString("1")))
	assert.Equal(t, 4, BytesFromString("123").Count(nilBytes))
}

func TestBytes_CountR(t *testing.T) {
	assert.Equal(t, 0, nilBytes.CountR([]byte{'0'}))
	assert.Equal(t, 0, BytesFromString("123").CountR([]byte{'0'}))
	assert.Equal(t, 1, BytesFromString("123").CountR([]byte{'1'}))
	assert.Equal(t, 2, BytesFromString("1231").CountR([]byte{'1'}))
	assert.Equal(t, 4, BytesFromString("123").CountR(nil))
}

func TestBytes_Equal(t *testing.T) {
	assert.True(t, nilBytes.Equal(emptyBytes))
	assert.True(t, BytesFromString("123").Equal(BytesFromString("123")))
	assert.False(t, BytesFromString("123").Equal(BytesFromString("456")))
}

func TestBytes_EqualFold(t *testing.T) {
	assert.True(t, nilBytes.EqualFold(emptyBytes))
	assert.True(t, BytesFromString("abc").EqualFold(BytesFromString("ABC")))
	assert.False(t, BytesFromString("123").EqualFold(BytesFromString("abc")))
}

func TestBytes_EqualFoldR(t *testing.T) {
	assert.True(t, nilBytes.EqualFoldR(nil))
	assert.True(t, BytesFromString("abc").EqualFoldR([]byte("ABC")))
	assert.False(t, BytesFromString("123").EqualFoldR([]byte("abc")))
}

func TestBytes_EqualR(t *testing.T) {
	assert.True(t, nilBytes.EqualR(nil))
	assert.True(t, BytesFromString("123").EqualR([]byte("123")))
	assert.False(t, BytesFromString("123").EqualR([]byte("456")))
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
	// todo
}

func TestBytes_MarshalText(t *testing.T) {
	// todo
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

func TestBytes_UnmarshalBinary(t *testing.T) {
	// todo
}

func TestBytes_UnmarshalText(t *testing.T) {
	// todo
}

func TestBytes_MarshalBSON(t *testing.T) {
	encoded, err := bson.Marshal(BytesFromString("1"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(encoded)
	var decoded Bytes
	if err = bson.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, BytesFromString("1"), decoded)
}

func TestNewBytes(t *testing.T) {
	// todo
}

func TestNewBytesSlice(t *testing.T) {
	// todo
}
