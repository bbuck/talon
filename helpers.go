package talon

import (
	"reflect"
	"time"
)

// convert an arbitrary struct into properties
func structToMap(i interface{}) Properties {
	value := reflect.ValueOf(i)
	typ := value.Type()
	props := make(Properties)
	if value.Kind() != reflect.Struct {
		return props
	}

	fc := value.NumField()
	for i := 0; i < fc; i++ {
		field := typ.Field(i)
		var key string
		if tag, ok := field.Tag.Lookup("talon"); ok {
			key = tag
		} else {
			key = field.Name
		}
		props[key] = value.Field(i).Interface()
	}

	return props
}

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
		bs, err := i.MarshalNeo()

		return string(bs), err
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := Uint(val.Uint())
		bs, err := u.MarshalNeo()

		return string(bs), err
	case reflect.String:
		s := String(val.String())
		bs, err := s.MarshalNeo()

		return string(bs), err
	}

	if t, ok := val.Interface().(time.Time); ok {
		nt := Time{t}
		bs, _ := nt.MarshalNeo()

		return string(bs), nil
	}

	return "", nil
}
