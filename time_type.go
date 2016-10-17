package talon

import "time"

const timeFormat = time.RFC3339

type Time struct {
	time.Time
}

func (t Time) MarshalTalon() ([]byte, error) {
	tstr := t.Format(timeFormat)

	return []byte(tstr), nil
}

func (t *Time) UnmarshalTalon(bs []byte) error {
	str := string(bs)
	pt, err := time.Parse(timeFormat, str)
	if err == nil {
		*t = Time{pt}
	}

	return err
}
