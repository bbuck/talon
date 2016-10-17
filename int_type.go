package talon

import "strconv"

type Int int64

// MarshalNeo defines a way for Int64 to turn into a neo value.
func (i Int) MarshalTalon() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(i), 10)), nil
}

func (i *Int) UnmarshalTalon(bs []byte) error {
	str := string(bs)
	i64, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		*i = Int(i64)
	}

	return err
}
