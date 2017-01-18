package talon

import (
	"reflect"
	"strings"
)

type loadObject struct {
	Type     reflect.Type
	Value    reflect.Value
	FieldMap map[string]reflect.Value
}

func newLoadObject(val reflect.Value) *loadObject {
	lo := &loadObject{
		Value:    val,
		Type:     val.Type(),
		FieldMap: make(map[string]reflect.Value),
	}

	for i, fc := 0, lo.Type.NumField(); i < fc; i++ {
		fv := lo.Value.Field(i)
		// we only consider fields we can set
		if !fv.CanSet() {
			continue
		}
		sf := lo.Type.Field(i)
		if sf.Anonymous {
			continue
		}
		if tagVal, ok := sf.Tag.Lookup("talon"); ok {
			lo.FieldMap[tagVal] = fv
			continue
		}
		lo.FieldMap[sf.Name] = fv
		lo.FieldMap[strings.ToLower(sf.Name)] = fv
	}

	return lo
}
