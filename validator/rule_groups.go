package validator

import (
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
	rules := []validation.Rule{
		LengthRule(lang, field, password.MinLength, password.MaxLength),
		FormatRule(lang, field, password.Regex),
	}

	if required {
		rules = append(rules, RequiredRule(lang, field))
	}

	return rules
}
