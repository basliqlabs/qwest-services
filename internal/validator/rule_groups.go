package validator

import (
	"fmt"

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
			return fmt.Errorf("password is not valid")
		}
		valid, err := password.IsValid(s)
		if !valid {
			return fmt.Errorf("password is not valid")
		}
		if err != nil {
			return err
		}
		return nil
	}

	rules := []validation.Rule{
		LengthRule(lang, field, password.MinLength, password.MaxLength),
		validation.By(isValid),
	}

	if required {
		rules = append(rules, RequiredRule(lang, field))
	}

	return rules
}
