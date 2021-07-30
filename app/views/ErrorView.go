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

func BodyParsingError() string {
	return fmt.Sprintf("Error parsing response body")
}
