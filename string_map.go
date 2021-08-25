package immutable

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
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

func (t StringMap) MarshalText() ([]byte, error) {
	return json.Marshal(t.m)
}

func (t *StringMap) UnmarshalText(bytes []byte) error {
	var m map[string]string
	if err := json.Unmarshal(bytes, &m); err != nil {
		return err
	}
	*t = StringMap{m}
	return nil
}

func (t StringMap) MarshalBinary() (data []byte, err error) {
	return t.MarshalBSON()
}

func (t *StringMap) UnmarshalBinary(data []byte) error {
	return t.UnmarshalBSON(data)
}

type stringMapDTO struct {
	Data map[string]string `bson:"data"`
}

func (t StringMap) MarshalBSON() ([]byte, error) {
	return bson.Marshal(stringMapDTO{Data: t.m})
}

func (t *StringMap) UnmarshalBSON(data []byte) error {
	var p stringMapDTO
	if err := bson.Unmarshal(data, &p); err != nil {
		return err
	}
	*t = StringMap{p.Data}
	return nil
}
