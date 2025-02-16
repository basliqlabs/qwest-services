package email

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	MinLength = 6
	MaxLength = 45
	Regex     = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

var (
	ErrMinLength     = fmt.Errorf("email must be at least %d characters long", MaxLength)
	ErrMaxLength     = fmt.Errorf("email must be at most %d characters long", MinLength)
	ErrRegexMismatch = fmt.Errorf("email failed to compile with regexp")
)

func IsValid(email string) (bool, error) {
	if len(email) < MinLength {
		return false, ErrMinLength
	}

	if len(email) > MaxLength {
		return false, ErrMaxLength
	}

	match, err := regexp.MatchString(Regex, strings.ToLower(email))
	if err != nil {
		return false, ErrRegexMismatch
	}
	return match, nil
}
