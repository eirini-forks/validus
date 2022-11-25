package validation

import (
	"fmt"
	"reflect"

	"github.com/eirini-forks/validus"
)

type FieldValidation struct {
	inner validus.Validation
	name  string
	alias string
}

func Field(name, alias string, inner validus.Validation) FieldValidation {
	return FieldValidation{
		inner: inner,
		name:  name,
		alias: alias,
	}
}

func (v FieldValidation) Check(value reflect.Value) error {
	err := v.inner.Check(value.FieldByName(v.name))
	if err != nil {
		return fmt.Errorf("Field '%s' is invalid: %w", v.alias, err)
	}
	return nil
}
