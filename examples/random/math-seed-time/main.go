package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	b := make([]byte, 4)
	_, err := rand.Read(b) // HL
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(b)
}
