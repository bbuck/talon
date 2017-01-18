package talon

import (
	"errors"
	"reflect"
)

// This files contains the code that loads query results into Go values, like
// talon.Cypher(...).Load(&Person{}) or talon.Cypher(...).Load(&Person{}, &[]Likes)
// It will (potentially) can an unclear jumble of relfection and dynamic object
// creation/mapping logic which may or may not be any fun at all.

// Load will process the cypher query and then attempt to load it's results
// into the given targets dynamically.
func (q *Query) Load(targets ...interface{}) error {
	if len(targets) == 0 {
		return errors.New("Load must be given targets to load results into.")
	}

	targetValues := getLoadObjects(targets)

	return nil
}

func getLoadObjects(ifaces []interface{}) ([]*loadObject, error) {
	targetValues := make([]*loadObject, len(ifaces))
	for idx, target := range ifaces {
		val := reflect.ValueOf(target)
		if !validateTarget(val) {
			return targetValues, errors.New("target values must be pointers")
		}
		targetValues[idx] = newLoadObject(val.Elem())
	}

	return targetValues, nil
}

func validateTarget(val reflect.Value) bool {
	return val.Kind() == reflect.Ptr
}

func fetchRowCount(targets []*loadObject) recordCount {
	return noRecords
}
