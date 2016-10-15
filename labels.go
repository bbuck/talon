// Copyright (c) 2016 Brandon Buck

package talon

import "bytes"

type Labels []string

func (l Labels) String() string {
	buf := new(bytes.Buffer)
	for _, lbl := range l {
		buf.WriteRune(':')
		buf.WriteString(lbl)
	}

	return buf.String()
}
