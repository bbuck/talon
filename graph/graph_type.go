package graph

const (
	NodeType GraphType = iota
	RelationshipType
)

type GraphTyper interface {
	Type() GraphType
}
