package injection

import (
	ut "github.com/go-playground/universal-translator"
	leafTranslatorValidator "github.com/paulusrobin/leaf-utilities/translator/validator"
)

type Translator struct {
	Validator *ut.UniversalTranslator
}

func NewValidatorTranslator() Translator {
	return Translator{
		Validator: leafTranslatorValidator.GetTranslator(),
	}
}
