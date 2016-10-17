// Copyright (c) 2016 Brandon Buck

package talon

import "reflect"

// convert an arbitrary Go type to a talon type.
func toTalonType(i interface{}) interface{} {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Int(val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Uint(val.Uint())
	case reflect.String:
		return String(val.String())
	case reflect.Float32, reflect.Float64:
		return Float(val.Float())
	}

	return i
}
