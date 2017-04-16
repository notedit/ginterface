package goterface

import (
	"errors"
	"reflect"
)

func Version() string {
	return "0.0.1"
}

type Goterface struct {
	data interface{}
}

func New(data interface{}) *Goterface {
	g := new(Goterface)
	g.data = data
	return g
}

func (g *Goterface) Interface() interface{} {
	return g.data
}

func (g *Goterface) Map() (map[string]interface{}, error) {
	if m, ok := (g.data).(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("type assertion to map[string]interface{} failed")
}

func (g *Goterface) Array() ([]*Goterface, error) {
	if arr, ok := (g.data).([]interface{}); ok {
		retArr := make([]*Goterface, 0, len(arr))
		for _, a := range arr {
			if a == nil {
				continue
			}
			s := &Goterface{a}
			retArr = append(retArr, s)
		}
		return retArr, nil
	}
	return nil, errors.New("type assertion to []interface{} failed")
}

func (g *Goterface) Set(key string, val interface{}) {
	m, err := g.Map()
	if err != nil {
		return
	}
	m[key] = val
}

func (g *Goterface) Del(key string) {
	m, err := g.Map()
	if err != nil {
		return
	}
	delete(m, key)
}

func (g *Goterface) Get(key string) *Goterface {
	m, err := g.Map()
	if err == nil {
		if val, ok := m[key]; ok {
			return &Goterface{val}
		}
	}
	return &Goterface{nil}
}

func (g *Goterface) CheckGet(key string) (*Goterface, bool) {
	m, err := g.Map()
	if err == nil {
		if val, ok := m[key]; ok {
			return &Goterface{val}, true
		}
	}
	return nil, false
}

func (g *Goterface) Bool() (bool, error) {
	if s, ok := (g.data).(bool); ok {
		return s, nil
	}
	return false, errors.New("type assertion to bool failed")
}

func (g *Goterface) String() (string, error) {
	if s, ok := (g.data).(string); ok {
		return s, nil
	}
	return "", errors.New("type assertion to string failed")
}

func (g *Goterface) Bytes() ([]byte, error) {
	if s, ok := (g.data).(string); ok {
		return []byte(s), nil
	}
	return nil, errors.New("type assertion to []byte failed")
}

func (g *Goterface) StringArray() ([]string, error) {
	arr, err := g.Array()
	if err != nil {
		return nil, err
	}
	retArr := make([]string, 0, len(arr))
	for _, a := range arr {
		if a == nil {
			retArr = append(retArr, "")
			continue
		}
		s, ok := a.String()
		if ok != nil {
			return nil, errors.New("type assertion to []string failed")
		}
		retArr = append(retArr, s)
	}
	return retArr, nil
}

func (g *Goterface) Int() (int, error) {
	switch g.data.(type) {
	case float32, float64:
		return int(reflect.ValueOf(g.data).Float()), nil
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(g.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(g.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

// Int64 coerces into an int64
func (g *Goterface) Int64() (int64, error) {
	switch g.data.(type) {
	case float32, float64:
		return int64(reflect.ValueOf(g.data).Float()), nil
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(g.data).Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(g.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

// Uint64 coerces into an uint64
func (g *Goterface) Uint64() (uint64, error) {
	switch g.data.(type) {
	case float32, float64:
		return uint64(reflect.ValueOf(g.data).Float()), nil
	case int, int8, int16, int32, int64:
		return uint64(reflect.ValueOf(g.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(g.data).Uint(), nil
	}
	return 0, errors.New("invalid value type")
}

func (g *Goterface) Float64() (float64, error) {
	switch g.data.(type) {
	case float32, float64:
		return reflect.ValueOf(g.data).Float(), nil
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(g.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(g.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}
