package tokenGenerator

import (
	"crypto/rand"
	"fmt"
)

const Amount = 16

func Generate() string {
	b := make([]byte, Amount)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}
