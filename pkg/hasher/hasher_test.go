package hasher

import (
	"fmt"
	"testing"
)

func TestCheckPasswordHash() {
	fmt.Println(CheckPasswordHash("password", "$2a$10$Kt0YB3SgXJuUpek5anTDguHyKXUEbE4EIyzQXrfzYzsNB9ExZflSe"))
}
