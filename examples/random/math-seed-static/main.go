package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1)

	b := make([]byte, 4)
	_, err := rand.Read(b) // HL
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(b)
}
