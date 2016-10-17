package talon

// String is used in place of the default Go string type so that methods can
// be added to it allowing it to match the Marshaler and Unmarshaler interfaces.
type String string

func (s String) String() string {
	return string(s)
}

// MarshalTalon brings String inline with the Marshaler interface for talon
// query building. This method will convert a string into a valid representaiton
// as a property value in a Cypher query.
func (s String) MarshalTalon() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

// UnmarshalTalon is designed to take a value from a Cypher result and convert
// it into a valid String type. This method _shouldn't_ be necessary as the
// bytes can be converted directly into a string (which is what this does) and
// should be done by the driver. But the implementation is here regardless for
// when it's necessary.
func (s *String) UnmarshalTalon(bs []byte) error {
	*s = String(bs)

	return nil
}
