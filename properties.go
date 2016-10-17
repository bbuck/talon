// Copyright (c) 2016 Brandon Buck

package talon

import (
	"bytes"
	"reflect"
	"time"
)

// Properties is a map[string]interface{} wrapper with a special string function
// designed to produce properties for Neo4j.
type Properties map[string]interface{}

// String brings Properties inline with fmt.Stringer and produced a Neo4j
// compatible propety map
func (p Properties) String() string {
	if len(p) == 0 {
		return ""
	}

	buf := new(bytes.Buffer)

	buf.WriteRune('{')

	buf.WriteRune('}')

	return buf.String()
}

// helpers

// merges a and b property maps, does not alter either input
func mergeProperties(old, new Properties) Properties {
	props := make(Properties)
	for key, val := range old {
		props[key] = val
	}

	for key, val := range new {
		props[key] = val
	}

	return props
}

func ifaceToNeoPropertyString(i interface{}) (string, error) {
	return valueNeoPropertyValueString(reflect.ValueOf(i))
}

func valueNeoPropertyValueString(val reflect.Value) (string, error) {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := Int(val.Int())
		bs, err := i.MarshalTalon()

		return string(bs), err
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := Uint(val.Uint())
		bs, err := u.MarshalTalon()

		return string(bs), err
	case reflect.String:
		s := String(val.String())
		bs, err := s.MarshalTalon()

		return string(bs), err
	}

	if t, ok := val.Interface().(time.Time); ok {
		nt := Time{t}
		bs, _ := nt.MarshalTalon()

		return string(bs), nil
	}

	return "", nil
}
