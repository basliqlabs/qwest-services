package username

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	MinUserNameLength = 6
	MaxUserNameLength = 32
)

var UserNameRegex = fmt.Sprintf("^[a-z]{1}[a-z0-9]{%d,%d}$", MinUserNameLength-1, MaxUserNameLength-1)

var (
	ErrMinLength     = fmt.Errorf("username must be lass than %d characters", MaxUserNameLength)
	ErrMaxLength     = fmt.Errorf("username must be greater than %d characters", MinUserNameLength)
	ErrRegexMismatch = fmt.Errorf("username failed to compile with regexp")
)

func IsValid(username string) (bool, error) {
	if len(username) < MinUserNameLength {
		return false, ErrMinLength
	}

	if len(username) > MaxUserNameLength {
		return false, ErrMaxLength
	}

	match, err := regexp.MatchString(UserNameRegex, strings.ToLower(username))
	if err != nil {
		return false, ErrRegexMismatch
	}
	return match, nil
}
