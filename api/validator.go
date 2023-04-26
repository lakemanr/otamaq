package api

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

var validName validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if name, ok := fieldLevel.Field().Interface().(string); ok {
		for _, r := range name {
			isValid := unicode.IsLetter(r) || unicode.IsSpace(r) || unicode.IsDigit(r) || r == '-' || r == '_'
			if !isValid {
				return false
			}
		}
		return true
	}
	return false
}
