package validus

import "reflect"

type Validation interface {
	Check(reflect.Value) error
}

func Validate(value interface{}, validation Validation) error {
	return validation.Check(reflect.ValueOf(value))
}
