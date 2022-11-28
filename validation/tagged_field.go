package validation

import (
	"fmt"
	"reflect"

	"github.com/eirini-forks/validus"
	"github.com/fatih/structtag"
)

type TaggedFieldValidation struct {
	inner validus.Validation
	name  string
	tag   string
}

func TaggedField(name, tag string, inner validus.Validation) TaggedFieldValidation {
	return TaggedFieldValidation{
		inner: inner,
		name:  name,
		tag:   tag,
	}
}

func JSONField(name string, inner validus.Validation) TaggedFieldValidation {
	return TaggedField(name, "json", inner)
}

func (v TaggedFieldValidation) Check(value reflect.Value) error {
	if !value.IsValid() {
		return fmt.Errorf("value is nil")
	}

	valueType := value.Type()
	if valueType.Kind() == reflect.Pointer {
		if value.IsNil() {
			return fmt.Errorf("value is nil")
		}

		valueType = valueType.Elem()
		value = reflect.Indirect(value)
	}

	if valueType.Kind() != reflect.Struct {
		return fmt.Errorf("value '%v' is not a struct", value)
	}

	field, ok := valueType.FieldByName(v.name)
	if !ok {
		return fmt.Errorf("no field '%s' in type '%v'", v.name, value.Type())
	}

	fieldValue := value.FieldByName(v.name)

	tags, err := structtag.Parse(string(field.Tag))
	if err != nil {
		return fmt.Errorf("failed to parse tags for '%s': %w", v.name, err)
	}

	tag, err := tags.Get(v.tag)
	if err != nil {
		return fmt.Errorf("failed to get tag '%s' for '%s': %w", v.tag, v.name, err)
	}

	err = v.inner.Check(fieldValue)
	if err != nil {
		return fmt.Errorf("Field '%s' is invalid: %w", tag.Name, err)
	}

	return nil
}
