package email

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	MinEmailLength = 6
	MaxEmailLength = 45
	EmailRegex     = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

var (
	ErrMinLength     = fmt.Errorf("email must be at least %d characters long", MaxEmailLength)
	ErrMaxLength     = fmt.Errorf("email must be at most %d characters long", MinEmailLength)
	ErrRegexMismatch = fmt.Errorf("email failed to compile with regexp")
)

func IsValid(email string) (bool, error) {
	if len(email) < MinEmailLength {
		return false, ErrMinLength
	}

	if len(email) > MaxEmailLength {
		return false, ErrMaxLength
	}

	match, err := regexp.MatchString(EmailRegex, strings.ToLower(email))
	if err != nil {
		return false, ErrRegexMismatch
	}
	return match, nil
}
