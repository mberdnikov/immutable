package immutable

import (
	"encoding/json"
)

type StringMap struct {
	m map[string]string
}

func NewStringMap(m map[string]string, copy bool) StringMap {
	if !copy {
		return StringMap{m}
	}
	if len(m) == 0 {
		return StringMap{}
	}
	cp := make(map[string]string, len(m))
	for key, value := range m {
		cp[key] = value
	}
	return StringMap{cp}
}

func (t StringMap) Len() int {
	return len(t.m)
}

func (t StringMap) Copy(m map[string]string) {
	for key, value := range t.m {
		m[key] = value
	}
}

func (t StringMap) Keys() []string {
	if len(t.m) == 0 {
		return nil
	}
	res := make([]string, 0, len(t.m))
	for key := range t.m {
		res = append(res, key)
	}
	return res
}

func (t StringMap) Get(key string) (string, bool) {
	if t.m == nil {
		return "", false
	}
	val, has := t.m[key]
	return val, has
}

func (t StringMap) Has(key string) bool {
	if t.m == nil {
		return false
	}
	_, has := t.m[key]
	return has
}

func (t StringMap) Default(key string, def string) string {
	if t.m == nil {
		return def
	}
	val, has := t.m[key]
	if !has {
		return def
	}
	return val
}

func (t StringMap) With(key, value string) StringMap {
	a := map[string]string{key: value}
	if len(t.m) == 0 {
		return StringMap{m: a}
	}
	return t.merge(a)
}

func (t StringMap) Merge(sm StringMap) StringMap {
	if len(t.m) == 0 {
		return sm
	}
	return t.merge(sm.m)
}

func (t StringMap) Without(key string) StringMap {
	if t.m == nil {
		return t
	}
	if _, has := t.m[key]; !has {
		return t
	}
	if len(t.m) == 1 {
		return StringMap{}
	}
	cp := make(map[string]string, len(t.m)-1)
	for k, v := range t.m {
		if k == key {
			continue
		}
		cp[k] = v
	}
	return StringMap{cp}
}

func (t StringMap) IsEqual(sm StringMap) bool {
	if len(t.m) != len(sm.m) {
		return false
	}
	for key, value1 := range t.m {
		value2, has := sm.m[key]
		if !has {
			return false
		}
		if value1 != value2 {
			return false
		}
	}
	return true
}

func (t StringMap) merge(sm map[string]string) StringMap {
	count := len(t.m)
	for key := range sm {
		if _, has := t.m[key]; has {
			count++
		}
	}
	result := make(map[string]string, count)
	for key, value := range t.m {
		result[key] = value
	}
	for key, value := range sm {
		result[key] = value
	}
	return StringMap{result}
}

func (t StringMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.m)
}

func (t *StringMap) UnmarshalJSON(bytes []byte) error {
	var m map[string]string
	if err := json.Unmarshal(bytes, &m); err != nil {
		return err
	}
	*t = StringMap{m}
	return nil
}

func (t StringMap) MarshalText() ([]byte, error) {
	return t.MarshalJSON()
}

func (t *StringMap) UnmarshalText(bytes []byte) error {
	return t.UnmarshalJSON(bytes)
}

func (t StringMap) MarshalBinary() (data []byte, err error) {
	size := 0
	size += varintSize(uint64(len(t.m)))
	for k, v := range t.m {
		size += varintSize(uint64(len(k))) + len(k)
		size += varintSize(uint64(len(v))) + len(v)
	}

	data = make([]byte, size)
	off := 0
	writeInt := func(i int) {
		n := putVarint(data[off:], uint64(i))
		off += n
	}
	writeStr := func(s string) {
		off += copy(data[off:], s)
	}

	writeInt(len(t.m))
	for k, v := range t.m {
		writeInt(len(k))
		writeStr(k)
		writeInt(len(v))
		writeStr(v)
	}
	return data, nil
}

func (t *StringMap) UnmarshalBinary(data []byte) error {
	off := 0
	readInt := func() int {
		res, n := getVarint(data[off:])
		off += n
		return int(res)
	}
	readStr := func(l int) string {
		res := string(data[off : off+l])
		off += l
		return res
	}

	l := readInt()
	m := make(map[string]string, l)
	for i := 0; i < l; i++ {
		kl := readInt()
		k := readStr(kl)
		vl := readInt()
		v := readStr(vl)
		m[k] = v
	}
	*t = StringMap{m}
	return nil
}
