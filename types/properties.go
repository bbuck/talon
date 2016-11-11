// Copyright (c) 2016 Brandon Buck

package types

import (
	"bytes"
	"reflect"
	"sort"
	"time"
)

const divider = "$$"

// Properties is a map[string]interface{} wrapper with a special string function
// designed to produce properties for Neo4j.
type Properties map[string]interface{}

// String brings Properties inline with fmt.Stringer and produced a Neo4j
// compatible propety map
func (p Properties) String() string {
	return p.StringWithPostfixedProperties("")
}

// StringWithPostfixedProperties returns the same property string as String
// except that all inject property names include the divider ($$) and a postfix
// value. This aids in sorting duplicates. So a node named 'A' would produce
// a key of 'one' as 'one$$node_a'.
func (p Properties) StringWithPostfixedProperties(postfix string) string {
	if len(p) == 0 {
		return ""
	}

	buf := new(bytes.Buffer)

	buf.WriteRune('{')
	keys := p.Keys()
	for i, key := range keys {
		buf.WriteString(key)
		buf.WriteString(": {")
		buf.WriteString(key)
		if len(postfix) > 0 {
			buf.WriteString(divider)
			buf.WriteString(postfix)
		}
		buf.WriteRune('}')
		if i != len(keys)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteRune('}')

	return buf.String()
}

// ForQuery returns the same key/value pairs just with a postfix applied to the
// keys as specificed. This is used when building strings and passing data
// through the driver to prevent property collisions throughout the query.
func (p Properties) ForQuery(postfix string) (Properties, error) {
	qprops := make(map[string]interface{})
	for key, val := range p {
		newKey := key
		if len(postfix) > 0 {
			newKey = key + divider + postfix
		}
		val, err := marshalTalonValue(val)
		if err != nil {
			return nil, err
		}
		qprops[newKey] = val
	}

	return qprops, nil
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

// Merge merges the current Properties key/value pairs with those of the given
// Properties object. This does not modify the current or other input objects
// it instead returns a new Property map representing the merged properties.
func (p Properties) Merge(other Properties) Properties {
	props := make(Properties)
	for key, val := range p {
		props[key] = val
	}

	for key, val := range other {
		props[key] = val
	}

	return props
}

func marshalTalonValue(i interface{}) (interface{}, error) {
	if tm, ok := i.(Marshaler); ok {
		bs, err := tm.MarshalTalon()
		if err != nil {
			return nil, err
		}

		return string(bs), nil
	}

	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Complex64, reflect.Complex128:
		c128 := val.Complex()
		c := Complex(c128)
		bs, err := c.MarshalTalon()
		if err != nil {
			return nil, err
		}

		return string(bs), nil
	}

	if t, ok := i.(time.Time); ok {
		tt := NewTime(t)
		bs, err := tt.MarshalTalon()
		if err != nil {
			return nil, err
		}

		return string(bs), nil
	}

	return i, nil
}
