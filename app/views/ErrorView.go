package views

import (
	"fmt"
)

func FieldNotFoundError(field string) string {
	return fmt.Sprintf(`This %s cannot be found`, field)
}

func FieldRequiredError(field string) string {
	return fmt.Sprintf(`The field '%s' is missing`, field)
}

func FieldForbiddenError(field string) string {
	return fmt.Sprintf(`You don't have the right to modify this %s`, field)
}

func BodyParsingError() string {
	return "Error parsing response body"
}
