// Copyright (c) 2016 Brandon Buck

package talon

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func toNeoType(i interface{}) interface{} {
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

// Int

type Int int64

// MarshalNeo defines a way for Int64 to turn into a neo value.
func (i Int) MarshalNeo() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(i), 10)), nil
}

func (Int) UnmarshalNeo(bs []byte) (interface{}, error) {
	str := string(bs)

	return strconv.ParseInt(str, 10, 64)
}

// Uint

type Uint uint64

func (u Uint) MarshalNeo() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(u), 10)), nil
}

func (Uint) Unmarshal(bs []byte) (interface{}, error) {
	str := string(bs)

	return strconv.ParseUint(str, 10, 64)
}

// Float

type Float float64

func (f Float) MarshalNeo() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(f), byte('f'), -1, 64)), nil
}

func (Float) UnmarshalNeo(bs []byte) (interface{}, error) {
	str := string(bs)

	return strconv.ParseFloat(str, 64)
}

// String

type String string

func (s String) MarshalNeo() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", string(s))), nil
}

func (String) UnmarshalNeo(bs []byte) (interface{}, error) {
	return string(bs), nil
}

// Time

const timeFormat = time.RFC3339

type Time struct {
	time.Time
}

func (t Time) MarshalNeo() ([]byte, error) {
	tstr := t.Format(timeFormat)

	return []byte(tstr), nil
}

func (Time) UnmarshalNeo(tbytes []byte) (interface{}, error) {
	str := string(tbytes)

	return time.Parse(timeFormat, str)
}
