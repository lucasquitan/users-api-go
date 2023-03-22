package validation

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/lucasquitan/users-api-go/src/configuration/rest_err"
)

var (
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError               // UnmarshalTypeErro is when a JSON value that was not appropriate for a value of a specific
	var jsonValidationError validator.ValidationErrors // ValidationErrors is an array of FieldError's for use in custom error messages post validation.

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError(fmt.Sprintf("Invalid field type on fied %s. Received %s type.", jsonErr.Field, jsonErr.Value))

	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("There are invalid fields", errorCauses)

	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
