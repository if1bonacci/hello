package main

import (
	"fmt"
	"log"

	"github.com/if1bonacci/lets-go-chat/pkg/hasher"
)

func main() {
	hasher, err := hasher.HashPassword("password")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hasher)
}
