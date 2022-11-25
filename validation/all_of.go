package validation

import (
	"reflect"

	"github.com/eirini-forks/validus"
)

type AllOfValidation struct {
	validations []validus.Validation
}

func AllOf(validations ...validus.Validation) AllOfValidation {
	return AllOfValidation{validations: validations}
}

func (v AllOfValidation) Check(value reflect.Value) error {
	for _, vn := range v.validations {
		err := vn.Check(value)
		if err != nil {
			return err
		}
	}
	return nil
}
