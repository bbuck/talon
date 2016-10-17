package talon

import "strconv"

type Float float64

func (f Float) MarshalTalon() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(f), byte('f'), -1, 64)), nil
}

func (f *Float) UnmarshalTalon(bs []byte) error {
	str := string(bs)
	f64, err := strconv.ParseFloat(str, 64)
	if err == nil {
		*f = Float(f64)
	}

	return err
}
