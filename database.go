// Copyright (c) 2016 Brandon Buck

package talon

import "github.com/bbuck/talon/types"

// DB represents a talon connection to a Neo4j database using Neo4j bolt behind
// the scenes.
type DB struct{}

// Connect will take the provided connection options and attempt to establish
// a connection to a Neo4j database.
func Connect(co ConnectOptions) (*DB, error) {
	return &DB{}, nil
}

func (d *DB) Cypher(q string, p types.Properties) *Query {
	return &Query{
		db:         d,
		rawCypher:  q,
		properties: p,
	}
}
