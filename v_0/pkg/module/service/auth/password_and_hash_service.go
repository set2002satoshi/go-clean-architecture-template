package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparisonPassAndHash(cp, p string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(cp), []byte(p)); err != nil {
		return true
	}
	return false
}

