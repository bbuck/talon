package talon

import bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"

// &{
//   metadata: map[
//     fields:[n]
//   ]
//   statement: 0xc420430000
//   closed: false
//   consumed: true
//   finishedConsume: false
//   pipelineIndex: 0
//   closeStatement: true
// }

type Rows struct {
	Columns []string

	boltRows bolt.Rows
}

type Row struct{}

func (r *Rows) Close() {
	r.boltRows.Close()
}

func (r *Rows) Next() (interface{}, error) {
	row, _, err := r.boltRows.NextNeo()

	return row, err
}

func (r *Rows) All() ([][]interface{}, map[string]interface{}, error) {
	return r.boltRows.All()
}
