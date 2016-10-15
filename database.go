// Copyright (c) 2016 Brandon Buck

package talon

// DB represents a talon connection to a Neo4j database using Neo4j bolt behind
// the scenes.
type DB struct{}

// ConnectOptions allows customiztaino of how to connect to a Neo4j database
// with talon.
type ConnectOptions struct{}

// Connect will take the provided connection options and attempt to establish
// a connection to a Neo4j database.
func Connect(co ConnectOptions) (*DB, error) {
	return &DB{}, nil
}
