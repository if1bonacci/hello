package main

import (
	"fmt"

	"github.com/if1bonacci/lets-go-chat/pkg/hasher/hasher"
)

func main() {
	hasher, err := hasher.HashPassword()
	fmt.Println("hello, world.")
}
