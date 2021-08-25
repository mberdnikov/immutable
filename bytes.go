package immutable

import (
	"bytes"
	"encoding/base64"
	"io"
	"unicode"
)

type Bytes struct {
	p []byte
}

func NewBytes(p []byte, copy bool) Bytes {
	if copy {
		p = append([]byte(nil), p...)
	}
	return Bytes{p}
}

func BytesFromString(s string) Bytes {
	return NewBytes([]byte(s), false)
}

func (t Bytes) Len() int {
	return len(t.p)
}

func (t Bytes) Slice(from, to int) Bytes {
	return NewBytes(t.p[from:to], false)
}

func (t Bytes) Copy(dst []byte) int {
	return copy(dst, t.p)
}

func (t Bytes) Append(list ...Bytes) Bytes {
	length := len(t.p)
	for _, item := range list {
		length += len(item.p)
	}
	result := make([]byte, length)
	offset := copy(result, t.p)
	for _, item := range list {
		offset += copy(result[offset:], item.p)
	}
	return Bytes{result}
}

func (t Bytes) Reader() io.Reader {
	return bytes.NewReader(t.p)
}

// Equal is equivalent to bytes.Equal
func (t Bytes) Equal(b Bytes) bool {
	return bytes.Equal(t.p, b.p)
}

// EqualR is equivalent to bytes.Equal
func (t Bytes) EqualR(b []byte) bool {
	return bytes.Equal(t.p, b)
}

// EqualFold is equivalent to bytes.EqualFold
func (t Bytes) EqualFold(s Bytes) bool {
	return bytes.EqualFold(t.p, s.p)
}

// EqualFoldR is equivalent to bytes.EqualFold
func (t Bytes) EqualFoldR(s []byte) bool {
	return bytes.EqualFold(t.p, s)
}

// Compare is equivalent to bytes.Compare
func (t Bytes) Compare(b Bytes) int {
	return bytes.Compare(t.p, b.p)
}

// CompareR is equivalent to bytes.Compare
func (t Bytes) CompareR(b []byte) int {
	return bytes.Compare(t.p, b)
}

// Count is equivalent to bytes.Count
func (t Bytes) Count(sep Bytes) int {
	return bytes.Count(t.p, sep.p)
}

// CountR is equivalent to bytes.Count
func (t Bytes) CountR(sep []byte) int {
	return bytes.Count(t.p, sep)
}

// Contains is equivalent to bytes.Contains
func (t Bytes) Contains(subslice Bytes) bool {
	return bytes.Contains(t.p, subslice.p)
}

// ContainsR is equivalent to bytes.Contains
func (t Bytes) ContainsR(subslice []byte) bool {
	return bytes.Contains(t.p, subslice)
}

// ContainsAny is equivalent to bytes.ContainsAny
func (t Bytes) ContainsAny(chars string) bool {
	return bytes.ContainsAny(t.p, chars)
}

// ContainsRune is equivalent to bytes.ContainsRune
func (t Bytes) ContainsRune(r rune) bool {
	return bytes.ContainsRune(t.p, r)
}

// Index is equivalent to bytes.Index
func (t Bytes) Index(sep Bytes) int {
	return bytes.Index(t.p, sep.p)
}

// IndexR is equivalent to bytes.Index
func (t Bytes) IndexR(sep []byte) int {
	return bytes.Index(t.p, sep)
}

// IndexAny is equivalent to bytes.IndexAny
func (t Bytes) IndexAny(chars string) int {
	return bytes.IndexAny(t.p, chars)
}

// IndexFunc is equivalent to bytes.IndexFunc
func (t Bytes) IndexFunc(f func(r rune) bool) int {
	return bytes.IndexFunc(t.p, f)
}

// IndexByte is equivalent to bytes.IndexByte
func (t Bytes) IndexByte(c byte) int {
	return bytes.IndexByte(t.p, c)
}

// IndexRune is equivalent to bytes.IndexRune
func (t Bytes) IndexRune(r rune) int {
	return bytes.IndexRune(t.p, r)
}

// LastIndex is equivalent to bytes.LastIndex
func (t Bytes) LastIndex(sep Bytes) int {
	return bytes.LastIndex(t.p, sep.p)
}

// LastIndexR is equivalent to bytes.LastIndex
func (t Bytes) LastIndexR(sep []byte) int {
	return bytes.LastIndex(t.p, sep)
}

// LastIndexAny is equivalent to bytes.LastIndexAny
func (t Bytes) LastIndexAny(chars string) int {
	return bytes.LastIndexAny(t.p, chars)
}

// LastIndexFunc is equivalent to bytes.LastIndexFunc
func (t Bytes) LastIndexFunc(f func(r rune) bool) int {
	return bytes.LastIndexFunc(t.p, f)
}

// LastIndexByte is equivalent to bytes.LastIndexByte
func (t Bytes) LastIndexByte(c byte) int {
	return bytes.LastIndexByte(t.p, c)
}

// Split is equivalent to bytes.Split
func (t Bytes) Split(sep Bytes) []Bytes {
	return bytesSlice(bytes.Split(t.p, sep.p))
}

// SplitR is equivalent to bytes.Split
func (t Bytes) SplitR(sep []byte) []Bytes {
	return bytesSlice(bytes.Split(t.p, sep))
}

// SplitN is equivalent to bytes.SplitN
func (t Bytes) SplitN(sep Bytes, n int) []Bytes {
	return bytesSlice(bytes.SplitN(t.p, sep.p, n))
}

// SplitNR is equivalent to bytes.SplitN
func (t Bytes) SplitNR(sep []byte, n int) []Bytes {
	return bytesSlice(bytes.SplitN(t.p, sep, n))
}

// SplitAfter is equivalent to bytes.SplitAfter
func (t Bytes) SplitAfter(sep Bytes) []Bytes {
	return bytesSlice(bytes.SplitAfter(t.p, sep.p))
}

// SplitAfterR is equivalent to bytes.SplitAfter
func (t Bytes) SplitAfterR(sep []byte) []Bytes {
	return bytesSlice(bytes.SplitAfter(t.p, sep))
}

// SplitAfterN is equivalent to bytes.SplitAfterN
func (t Bytes) SplitAfterN(sep Bytes, n int) []Bytes {
	return bytesSlice(bytes.SplitAfterN(t.p, sep.p, n))
}

// SplitAfterNR is equivalent to bytes.SplitAfterN
func (t Bytes) SplitAfterNR(sep []byte, n int) []Bytes {
	return bytesSlice(bytes.SplitAfterN(t.p, sep, n))
}

// Fields is equivalent to bytes.Fields
func (t Bytes) Fields() []Bytes {
	return bytesSlice(bytes.Fields(t.p))
}

// FieldsFunc is equivalent to bytes.FieldsFunc
func (t Bytes) FieldsFunc(f func(rune) bool) []Bytes {
	return bytesSlice(bytes.FieldsFunc(t.p, f))
}

// HasPrefix is equivalent to bytes.HasPrefix
func (t Bytes) HasPrefix(prefix Bytes) bool {
	return bytes.HasPrefix(t.p, prefix.p)
}

// HasPrefixR is equivalent to bytes.HasPrefix
func (t Bytes) HasPrefixR(prefix []byte) bool {
	return bytes.HasPrefix(t.p, prefix)
}

// HasSuffix is equivalent to bytes.HasSuffix
func (t Bytes) HasSuffix(suffix Bytes) bool {
	return bytes.HasSuffix(t.p, suffix.p)
}

// HasSuffixR is equivalent to bytes.HasSuffix
func (t Bytes) HasSuffixR(suffix []byte) bool {
	return bytes.HasSuffix(t.p, suffix)
}

// TrimPrefix is equivalent to bytes.TrimPrefix
func (t Bytes) TrimPrefix(prefix Bytes) Bytes {
	return NewBytes(bytes.TrimPrefix(t.p, prefix.p), false)
}

// TrimPrefixR is equivalent to bytes.TrimPrefix
func (t Bytes) TrimPrefixR(prefix []byte) Bytes {
	return NewBytes(bytes.TrimPrefix(t.p, prefix), false)
}

// TrimSuffix is equivalent to bytes.TrimSuffix
func (t Bytes) TrimSuffix(suffix Bytes) Bytes {
	return NewBytes(bytes.TrimSuffix(t.p, suffix.p), false)
}

// TrimSuffixR is equivalent to bytes.TrimSuffix
func (t Bytes) TrimSuffixR(suffix []byte) Bytes {
	return NewBytes(bytes.TrimSuffix(t.p, suffix), false)
}

// Map is equivalent to bytes.Map
func (t Bytes) Map(mapping func(r rune) rune) Bytes {
	return NewBytes(bytes.Map(mapping, t.p), false)
}

// Repeat is equivalent to bytes.Repeat
func (t Bytes) Repeat(count int) Bytes {
	return NewBytes(bytes.Repeat(t.p, count), false)
}

// ToUpper is equivalent to bytes.ToUpper
func (t Bytes) ToUpper() Bytes {
	return NewBytes(bytes.ToUpper(t.p), false)
}

// ToLower is equivalent to bytes.ToLower
func (t Bytes) ToLower() Bytes {
	return NewBytes(bytes.ToLower(t.p), false)
}

// ToTitle is equivalent to bytes.ToTitle
func (t Bytes) ToTitle() Bytes {
	return NewBytes(bytes.ToTitle(t.p), false)
}

// Title is equivalent to bytes.Title
func (t Bytes) Title() Bytes {
	return NewBytes(bytes.Title(t.p), false)
}

// ToUpperSpecial is equivalent to bytes.ToUpperSpecial
func (t Bytes) ToUpperSpecial(c unicode.SpecialCase) Bytes {
	return NewBytes(bytes.ToUpperSpecial(c, t.p), false)
}

// ToLowerSpecial is equivalent to bytes.ToLowerSpecial
func (t Bytes) ToLowerSpecial(c unicode.SpecialCase) Bytes {
	return NewBytes(bytes.ToLowerSpecial(c, t.p), false)
}

// ToTitleSpecial is equivalent to bytes.ToTitleSpecial
func (t Bytes) ToTitleSpecial(c unicode.SpecialCase) Bytes {
	return NewBytes(bytes.ToTitleSpecial(c, t.p), false)
}

// Trim is equivalent to bytes.Trim
func (t Bytes) Trim(cutset string) Bytes {
	return NewBytes(bytes.Trim(t.p, cutset), false)
}

// TrimLeft is equivalent to bytes.TrimLeft
func (t Bytes) TrimLeft(cutset string) Bytes {
	return NewBytes(bytes.TrimLeft(t.p, cutset), false)
}

// TrimRight is equivalent to bytes.TrimRight
func (t Bytes) TrimRight(cutset string) Bytes {
	return NewBytes(bytes.TrimRight(t.p, cutset), false)
}

// TrimFunc is equivalent to bytes.TrimFunc
func (t Bytes) TrimFunc(f func(r rune) bool) Bytes {
	return NewBytes(bytes.TrimFunc(t.p, f), false)
}

// TrimLeftFunc is equivalent to bytes.TrimLeftFunc
func (t Bytes) TrimLeftFunc(f func(r rune) bool) Bytes {
	return NewBytes(bytes.TrimLeftFunc(t.p, f), false)
}

// TrimRightFunc is equivalent to bytes.TrimRightFunc
func (t Bytes) TrimRightFunc(f func(r rune) bool) Bytes {
	return NewBytes(bytes.TrimRightFunc(t.p, f), false)
}

// TrimSpace is equivalent to bytes.TrimSpace
func (t Bytes) TrimSpace() Bytes {
	return NewBytes(bytes.TrimSpace(t.p), false)
}

// ToValidUTF8 is equivalent to bytes.ToValidUTF8
func (t Bytes) ToValidUTF8(replacement Bytes) Bytes {
	return NewBytes(bytes.ToValidUTF8(t.p, replacement.p), false)
}

// ToValidUTF8R is equivalent to bytes.ToValidUTF8
func (t Bytes) ToValidUTF8R(replacement []byte) Bytes {
	return NewBytes(bytes.ToValidUTF8(t.p, replacement), false)
}

// Runes is equivalent to bytes.Runes
func (t Bytes) Runes() []rune {
	return bytes.Runes(t.p)
}

// Replace is equivalent to bytes.Replace
func (t Bytes) Replace(old, new Bytes, n int) Bytes {
	return NewBytes(bytes.Replace(t.p, old.p, new.p, n), false)
}

// ReplaceR is equivalent to bytes.Replace
func (t Bytes) ReplaceR(old, new []byte, n int) Bytes {
	return NewBytes(bytes.Replace(t.p, old, new, n), false)
}

// ReplaceAll is equivalent to bytes.ReplaceAll
func (t Bytes) ReplaceAll(old, new Bytes) Bytes {
	return NewBytes(bytes.ReplaceAll(t.p, old.p, new.p), false)
}

// ReplaceAllR is equivalent to bytes.ReplaceAll
func (t Bytes) ReplaceAllR(old, new []byte) Bytes {
	return NewBytes(bytes.ReplaceAll(t.p, old, new), false)
}

func (t Bytes) MarshalText() (text []byte, err error) {
	buf := bytes.Buffer{}
	enc := base64.NewEncoder(base64.RawStdEncoding, &buf)
	_, err = enc.Write(t.p)
	return buf.Bytes(), err
}

func (t *Bytes) UnmarshalText(text []byte) error {
	r := bytes.NewReader(text)
	buf := bytes.Buffer{}
	dec := base64.NewDecoder(base64.RawStdEncoding, r)
	if _, err := io.Copy(&buf, dec); err != nil {
		return err
	}
	*t = Bytes{buf.Bytes()}
	return nil
}

func (t Bytes) MarshalBinary() (data []byte, err error) {
	return append([]byte(nil), t.p...), nil
}

func (t *Bytes) UnmarshalBinary(data []byte) error {
	*t = Bytes{append([]byte(nil), data...)}
	return nil
}

func bytesSlice(p [][]byte) []Bytes {
	res := make([]Bytes, len(p))
	for idx, i := range p {
		res[idx] = NewBytes(i, false)
	}
	return res
}

func JoinBytes(sep Bytes, parts ...Bytes) Bytes {
	return JoinBytesR(sep.p, parts...)
}

func JoinBytesR(sep []byte, parts ...Bytes) Bytes {
	if len(parts) == 0 {
		return NewBytes(nil, false)
	}
	if len(parts) == 1 {
		return parts[0]
	}

	n := len(sep) * (len(parts) - 1)
	for _, v := range parts {
		n += len(v.p)
	}

	b := make([]byte, n)
	bp := copy(b, parts[0].p)
	for _, v := range parts[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], v.p)
	}
	return NewBytes(b, false)
}
