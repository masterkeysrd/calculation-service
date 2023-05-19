package validator

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var defaultTranslations = Translations{
	"required": "{0} field is required",
	"email":    "{0} field is not a valid email address",
}

type Translations map[string]string

func translateError(err error, trans ut.Translator) error {
	if err == nil {
		return nil
	}

	validationErrors := newValidationErrors(err)

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		validationErrors.Add(e.Field(), translatedErr.Error())
	}

	return validationErrors
}
