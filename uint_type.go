package talon

import "strconv"

type Uint uint64

func (u Uint) MarshalTalon() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(u), 10)), nil
}

func (u *Uint) UnmarshalTalon(bs []byte) error {
	str := string(bs)
	u64, err := strconv.ParseUint(str, 10, 64)
	if err == nil {
		*u = Uint(u64)
	}

	return err
}
