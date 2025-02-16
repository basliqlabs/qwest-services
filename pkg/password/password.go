package password

import (
	"fmt"
	"unicode"
)

const MinLength = 8
const MaxLength = 32

var (
	ErrMinLength     = fmt.Errorf("password must be at least %d characters long", MaxLength)
	ErrMaxLength     = fmt.Errorf("password must be at most %d characters long", MinLength)
	ErrInvalidFormat = fmt.Errorf("password has invalid format")
)

func IsValid(password string) (bool, error) {
	runes := len([]rune(password))
	if runes < MinLength {
		return false, ErrMinLength
	}

	if runes > MaxLength {
		return false, ErrMaxLength
	}

	number := false
	letter := false
	ascii := true

	for _, c := range password {
		switch {
		case c > unicode.MaxASCII:
			ascii = false
		case unicode.IsNumber(c):
			number = true
		case unicode.IsLetter(c):
			letter = true
		}
	}

	if number && letter && ascii {
		return true, nil
	}

	// TODO - add better error message
	return false, ErrInvalidFormat
}
