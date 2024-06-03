package helper

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	models "goAPI/models"
)

type Response struct {
	Status   int         `json:"status"`
	Messages string      `json:"messages"`
	Errors   interface{} `json:"errors"`
}

// ValidateUser validates the user struct
func ValidateUser(user *models.Users) map[string]string {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		errorList := make(map[string]string)

		for _, e := range errors {
			var errMsg string
			field, _ := reflect.TypeOf(*user).FieldByName(e.StructField())
			fieldName := field.Tag.Get("json")

			switch e.Tag() {
			case "required":
				errMsg = fmt.Sprintf("%s is required", fieldName)
			case "email":
				errMsg = fmt.Sprintf("%s is an invalid email", fieldName)
			case "min":
				errMsg = fmt.Sprintf("%s is too short, minimum length is %s", fieldName, e.Param())
			default:
				errMsg = fmt.Sprintf("%s is invalid", fieldName)
			}
			errorList[fieldName] = errMsg
		}
		return errorList
	}
	return nil
}
