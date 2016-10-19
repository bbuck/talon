// Copyright (c) 2016 Brandon Buck

package talon

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"
)

const divider = "$$ref"

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
	keys := p.Keys()
	for i, key := range keys {
		propKey := key
		if strings.Contains(key, divider) {
			keyParts := strings.Split(key, divider)
			key = keyParts[0]
		}
		buf.WriteString(key)
		buf.WriteString(": {")
		buf.WriteString(propKey)
		buf.WriteRune('}')
		if i != len(keys)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteRune('}')

	return buf.String()
}

// Keys returns an array of string values representing the keys in the map.
func (p Properties) Keys() []string {
	keys := make([]string, len(p))
	i := 0
	for key := range p {
		keys[i] = key
		i++
	}
	sort.Strings(keys)

	return keys
}

// MergeProperties will merge a and b property maps, does not alter either input
// maps.
func MergeProperties(old, new Properties) Properties {
	props := make(Properties)
	for key, val := range old {
		props[key] = val
	}

	for key, val := range new {
		if _, ok := props[key]; ok {
			for i := 1; true; i++ {
				newKey := fmt.Sprintf("%s$$ref%d", key, i)
				if _, ok := props[newKey]; !ok {
					key = newKey
					break
				}
			}
		}
		props[key] = val
	}

	return props
}

func ifaceToNeoPropertyString(i interface{}) (string, error) {
	return valueNeoPropertyValueString(reflect.ValueOf(i))
}

func valueNeoPropertyValueString(val reflect.Value) (string, error) {
	if t, ok := val.Interface().(time.Time); ok {
		nt := Time{t}
		bs, _ := nt.MarshalTalon()

		return string(bs), nil
	}

	return "", nil
}
