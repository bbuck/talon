// Copyright (c) 2016 Brandon Buck

package talon

import (
	"bytes"

	"github.com/bbuck/talon/types"
)

type NodeBuilder struct {
	Name       string
	Labels     Labels
	Properties types.Properties
}

// Node returns an empty NodeBuilder object ready to build a node used in a
// Cypher query.
func Node() *NodeBuilder {
	return new(NodeBuilder)
}

// Named assigns a name to the NodeBuilder object.
func (n *NodeBuilder) Named(name string) *NodeBuilder {
	n.Name = name

	return n
}

// Labeled applies a label or list of labels to the Node.
func (n *NodeBuilder) Labeled(lbls ...string) *NodeBuilder {
	for _, lbl := range lbls {
		n.Labels = append(n.Labels, lbl)
	}

	return n
}

// String brings NodeBuilder inline with the fmt.Stringer interface.
func (n *NodeBuilder) String() string {
	buf := new(bytes.Buffer)
	buf.WriteRune('(')
	if n.Name != "" {
		buf.WriteString(n.Name)
	}
	if len(n.Labels) > 0 {
		buf.WriteString(n.Labels.String())
	}
	if pstr := n.Properties.String(); pstr != "" {
		buf.WriteRune(' ')
		buf.WriteString(pstr)
	}
	buf.WriteRune(')')

	return buf.String()
}
