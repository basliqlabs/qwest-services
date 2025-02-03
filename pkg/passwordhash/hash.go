package passwordhash

import (
	"github.com/basliqlabs/qwest-services/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	const op = "passwordhash.Hash"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", richerror.New(op).WithKind(richerror.KindUnexpected).WithError(err)
	}
	return string(hash), nil
}

func Compare(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
