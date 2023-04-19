package hasher

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

var password = "password"

func TestHashPassword(t *testing.T) {
	hash, _ := HashPassword(password)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		t.Errorf("Hash Password is wrong")
	}
}
