package talon

import "time"

const DefaultTimeFormat = time.RFC3339Nano

type Time struct {
	time.Time
	OutputFormat string
}

func NewTime(t time.Time) Time {
	return Time{
		Time:         t,
		OutputFormat: DefaultTimeFormat,
	}
}

func NewTimeWithFormat(t time.Time, f string) Time {
	tt := NewTime(t)
	tt.OutputFormat = f

	return tt
}

func (t Time) MarshalTalon() ([]byte, error) {
	tstr := t.Format(t.OutputFormat)

	return []byte(tstr), nil
}

func (t *Time) UnmarshalTalon(bs []byte) error {
	str := string(bs)
	pt, err := time.Parse(t.OutputFormat, str)
	if err == nil {
		t.Time = pt
	}

	return err
}
