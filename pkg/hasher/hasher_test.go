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

func TestCheckPasswordHash(t *testing.T) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	status := CheckPasswordHash(password, string(hash))

	if !status {
		t.Errorf("Password is wrong")
	}
}

func TestFailCheckPasswordHash(t *testing.T) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	status := CheckPasswordHash("password2", string(hash))

	if status {
		t.Errorf("Password is wrong")
	}
}
