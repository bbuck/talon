package talon

import "github.com/bbuck/talon/types"

// Query reprsents a Talon query before it's been converted in Cypher
type Query struct {
	db         *DB
	rawCypher  string
	properties types.Properties
}
