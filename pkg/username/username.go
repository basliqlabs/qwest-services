package username

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const (
	MinUserNameLength = 6
	MaxUserNameLength = 32
)

var UserNameRegex = fmt.Sprintf("^[a-z]{1}[a-z0-9]{%d,%d}$", MinUserNameLength-1, MaxUserNameLength-1)

var (
	ErrMinLength     = fmt.Errorf("username must be at least %d characters long", MaxUserNameLength)
	ErrMaxLength     = fmt.Errorf("username must be at most %d characters long", MinUserNameLength)
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

func GenerateUnique(base string) string {
	base = strings.ToLower(base)
	reg := regexp.MustCompile("[^a-z0-9]+")
	base = reg.ReplaceAllString(base, "_")

	if len(base) < MinUserNameLength {
		base = base + strings.Repeat("0", MinUserNameLength-len(base))
	}

	if len(base) > MaxUserNameLength-6 {
		base = base[:MaxUserNameLength-6]
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	suffix := r.Intn(999999)

	return fmt.Sprintf("%s%06d", base, suffix)
}
