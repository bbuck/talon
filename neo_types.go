package talon

import (
	"fmt"
	"strconv"
	"time"
)

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
