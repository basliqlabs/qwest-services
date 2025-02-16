package validator

import (
	"errors"
	"fmt"
	"github.com/basliqlabs/qwest-services/pkg/translation"

	"github.com/basliqlabs/qwest-services/pkg/email"
	"github.com/basliqlabs/qwest-services/pkg/password"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func EmailRule(lang string, field string, required bool) []validation.Rule {
	rules := []validation.Rule{
		LengthRule(lang, field, email.MinLength, email.MaxLength),
		FormatRule(lang, field, email.Regex),
	}

	if required {
		rules = append(rules, RequiredRule(lang, field))
	}

	return rules
}

func PasswordRule(lang string, field string, required bool) []validation.Rule {
	isValid := func(value interface{}) error {
		s, ok := value.(string)
		if !ok {
			return fmt.Errorf(translation.TD(lang, "validation.invalid", map[string]any{
				"Field": translation.T(lang, field),
			}))
		}
		_, err := password.IsValid(s)

		if err != nil {
			switch {
			case errors.Is(err, password.ErrInvalidFormat):
				return fmt.Errorf(translation.TD(lang, "validation.invalid_password_format", map[string]any{
					"Field": translation.T(lang, field),
				}))
			case errors.Is(err, password.ErrMinLength):
				return fmt.Errorf(translation.TD(lang, "validation.min_length", map[string]any{
					"Field":     translation.T(lang, field),
					"MinLength": password.MinLength,
				}))
			case errors.Is(err, password.ErrMaxLength):
				return fmt.Errorf(translation.TD(lang, "validation.max_length", map[string]any{
					"Field":     translation.T(lang, field),
					"MaxLength": password.MaxLength,
				}))
			}
		}
		return nil
	}

	rules := []validation.Rule{
		validation.By(isValid),
	}

	if required {
		rules = append(rules, RequiredRule(lang, field))
	}

	return rules
}
