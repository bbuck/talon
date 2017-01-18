package talon

type recordCount int8

const (
	noRecords recordCount = iota
	oneRecord
	allRecords
)
