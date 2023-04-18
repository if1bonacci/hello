package hasher

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

var password = "password"

func TestCheckWrongPasswordHash(t *testing.T) {
	hash, _ := HashPassword(password)

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		t.Errorf("Hash Password is wrong")
	}

	// fmt.Println(CheckPasswordHash("password", "$2a$10$Kt0YB3SgXJuUpek5anTDguHyKXUEbE4EIyzQXrfzYzsNB9ExZflSe"))
}
