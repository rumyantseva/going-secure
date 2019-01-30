package main

import (
	"crypto/rand" // HL
	"fmt"
)

func main() {
	b := make([]byte, 4)

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println(b)
}
