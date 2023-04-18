package main

import (
	"fmt"

	"github.com/if1bonacci/lets-go-chat/pkg/hasher"
)

func main() {
	hasher, err := hasher.HashPassword("password")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(hasher)
}
