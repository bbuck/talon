// Copyright (c) 2016 Brandon Buck

package talon

import bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"

// ResultStats are some details about the results of the query, such as the
// number of of nodes, labels and properties added/set.
type ResultStats struct {
	LabelsAdded   int64
	NodesCreated  int64
	PropertiesSet int64
}

// Result represents a return from a non-row based query like a create/delete
// or upate where you're not using something like "return" in the query.
type Result struct {
	Stats ResultStats
	Type  string
}

// Close exits primarly to match Rows return behavior. When you run `Query`
// you have to close the rows object yourself. This is just here to prevent
// breakages.
func Close() { /* noop */ }

func wrapBoltResult(r bolt.Result) *Result {
	md := r.Metadata()
	stats := md["stats"].(map[string]interface{})
	return &Result{
		Stats: ResultStats{
			LabelsAdded:   stats["labels-added"].(int64),
			NodesCreated:  stats["nodes-created"].(int64),
			PropertiesSet: stats["properties-set"].(int64),
		},
		Type: md["type"].(string),
	}
}
