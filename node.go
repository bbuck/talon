// Copyright (c) 2016 Brandon Buck

package talon

import "bytes"

type NodeBuilder struct {
	name   string
	labels Labels
}

// Node returns an empty NodeBuilder object ready to build a node used in a
// Cypher query.
func Node() *NodeBuilder {
	return new(NodeBuilder)
}

// Named assigns a name to the NodeBuilder object.
func (n *NodeBuilder) Named(name string) *NodeBuilder {
	n.name = name

	return n
}

// Labeled applies a label or list of labels to the Node.
func (n *NodeBuilder) Labeled(lbls ...string) *NodeBuilder {
	for _, lbl := range lbls {
		n.labels = append(n.labels, lbl)
	}

	return n
}

// String brings NodeBuilder inline with the fmt.Stringer interface.
func (n *NodeBuilder) String() string {
	buf := new(bytes.Buffer)
	buf.WriteRune('(')
	if n.name != "" {
		buf.WriteString(n.name)
	}
	if len(n.labels) > 0 {
		buf.WriteString(n.labels.String())
	}
	buf.WriteRune(')')

	return buf.String()
}
