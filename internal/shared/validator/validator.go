package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	// Resolve field names from their JSON tag so validation errors reference
	// the wire field (e.g. "category_id") instead of the Go field name.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Validate runs struct validation and returns a map of JSON field name -> error
// message. Returns nil when the struct is valid.
func Validate(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return map[string]string{"unknown": err.Error()}
	}

	result := make(map[string]string, len(errs))
	for _, e := range errs {
		result[e.Field()] = formatMessage(e)
	}
	return result
}

// formatMessage converts a validator.FieldError into a user-friendly Indonesian
// error message. Unknown tags fall back to a generic message.
func formatMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "Field ini wajib diisi"
	case "email":
		return "Format email tidak valid"
	case "min":
		return "Minimal " + e.Param() + " karakter"
	case "max":
		return "Maksimal " + e.Param() + " karakter"
	case "uuid", "uuid4":
		return "Format UUID tidak valid"
	case "oneof":
		return "Nilai harus salah satu dari: " + e.Param()
	case "gt":
		return "Harus lebih besar dari " + e.Param()
	case "gte":
		return "Harus lebih besar atau sama dengan " + e.Param()
	case "lte":
		return "Harus lebih kecil atau sama dengan " + e.Param()
	default:
		return "Nilai tidak valid"
	}
}
