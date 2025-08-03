package validation

import (
	"errors"
	"reflect"
	"strings"
)

func Check(str string) bool {
	return str != ""
}

// ValidateStruct проверяет поля структуры по тегу `validate:"required"`
func ValidateStruct(s interface{}) error {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("validate")
		if strings.Contains(tag, "required") {
			value := val.Field(i)
			zero := reflect.Zero(field.Type).Interface()
			if reflect.DeepEqual(value.Interface(), zero) {
				return errors.New("field '" + field.Name + "' is required")
			}
		}
	}
	return nil
}
