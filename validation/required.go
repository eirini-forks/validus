package validation

import (
	"fmt"
	"reflect"
)

type RequiredValidation struct{}

func Required() RequiredValidation {
	return RequiredValidation{}
}

func (v RequiredValidation) Check(value reflect.Value) error {
	if value.IsZero() {
		return fmt.Errorf("a value is required, '%s' is a zero value", value)
	}
	return nil
}
