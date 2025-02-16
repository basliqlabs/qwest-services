package validator

import (
	"regexp"

	"github.com/basliqlabs/qwest-services/pkg/translation"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func RequiredRule(lang string, field string) validation.Rule {
	return validation.Required.Error(translation.TD(lang, "validation.required", map[string]any{
		"Field": translation.T(lang, field),
	}))
}

func LengthRule(lang string, field string, min int, max int) validation.Rule {
	return validation.Length(min, max).
		Error(translation.TD(lang, "validation.length", map[string]any{
			"Field":     translation.T(lang, field),
			"MinLength": min,
			"MaxLength": max,
		}))
}

func FormatRule(lang string, field string, regex string) validation.Rule {
	return validation.Match(regexp.MustCompile(regex)).Error(translation.TD(lang, "validation.invalid_format", map[string]any{
		"Field": translation.T(lang, field),
	}))
}
