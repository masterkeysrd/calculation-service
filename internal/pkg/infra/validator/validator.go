package validator

import (
	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
}

func NewValidator() *Validator {
	english := en.New()
	uni := ut.New(english, english)

	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	return &Validator{
		validate: validate,
		uni:      uni,
		trans:    trans,
	}
}

func (v *Validator) RegisterDefaultTranslations() {
	for tag, message := range defaultTranslations {
		v.RegisterTranslation(tag, message)
	}
}

func (v *Validator) RegisterTranslation(tag, message string) {
	v.validate.RegisterTranslation(tag, v.trans, func(ut ut.Translator) error {
		return ut.Add(tag, message, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())

		return t
	})
}

func (v *Validator) Validate(i interface{}) error {
	err := v.validate.Struct(i)
	errs := translateError(err, v.trans)

	return errs
}
