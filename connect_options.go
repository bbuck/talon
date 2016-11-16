package talon

import (
	"bytes"
	"fmt"
)

// ConnectOptions allows customiztaino of how to connect to a Neo4j database
// with talon.
type ConnectOptions struct {
	User string
	Pass string
	Host string
	Port uint16
}

// URL takes the options set for connection and generates a bolt connection
// string.
func (co *ConnectOptions) URL() string {
	buf := new(bytes.Buffer)
	buf.WriteString("bolt://")
	if co.User != "" {
		buf.WriteString(co.User)
		if co.Pass != "" {
			buf.WriteRune(':')
			buf.WriteString(co.Pass)
		}
		buf.WriteRune('@')
	}
	buf.WriteString(co.Host)
	if co.Port > 0 {
		buf.WriteRune(':')
		buf.WriteString(fmt.Sprintf("%d", co.Port))
	}

	return buf.String()
}
