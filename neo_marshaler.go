package talon

// NeoMarshaler defines functions that would turn a Go type into a valid neo
// marshaler.
type NeoMarshaler interface {
	MarshalNeo() ([]byte, error)
}

// NeoUnmarshaler defines functions needed to take a neo property and convert
// it into a Go type.
type NeoUnmarshaler interface {
	UnmarshalNeo(string) (interface{}, error)
}
