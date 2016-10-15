package talon

import "bytes"

// Properties is a map[string]interface{} wrapper with a special string function
// designed to produce properties for Neo4j.
type Properties map[string]interface{}

// String brings Properties inline with fmt.Stringer and produced a Neo4j
// compatible propety map
func (p Properties) String() string {
	buf := new(bytes.Buffer)

	buf.WriteRune('{')

	buf.WriteRune('}')

	return buf.String()
}
