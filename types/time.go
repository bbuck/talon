// Copyright (c) 2016 Brandon Buck

package types

import "time"

const DefaultTimeFormat = time.RFC3339Nano

// Time wraps the standard lib time.Time to provide marshaling capability with
// the Talon ORM. This includes a format string for the time.
type Time struct {
	time.Time
	OutputFormat string
}

// NewTime wraps the given time.Time in a talon Time object.
func NewTime(t time.Time) Time {
	return Time{
		Time:         t,
		OutputFormat: DefaultTimeFormat,
	}
}

// NewTimeWithFormat performs the same operation as NewTime but assigns a custom
// output format to the struct.
func NewTimeWithFormat(t time.Time, f string) Time {
	tt := NewTime(t)
	tt.OutputFormat = f

	return tt
}

// MarshalTalon allows Time to implement the Talon Marshaler interface.
func (t Time) MarshalTalon() ([]byte, error) {
	tstr := t.Format(t.OutputFormat)

	return []byte(tstr), nil
}

// UnmarshalTalon allows Time to implement the Talan Unmarshaler interface.
func (t *Time) UnmarshalTalon(bs []byte) error {
	str := string(bs)
	pt, err := time.Parse(t.OutputFormat, str)
	if err == nil {
		t.Time = pt
	}

	return err
}
