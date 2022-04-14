package injection

import (
	leafValidatorV10 "github.com/paulusrobin/leaf-utilities/validator/integrations/v10"
	leafValidator "github.com/paulusrobin/leaf-utilities/validator/validator"
)

func NewValidator(translator Translator) (leafValidator.Validator, error) {
	return leafValidatorV10.New(leafValidatorV10.WithTranslator(translator.Validator))
}
