package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		intn := rand.Intn(10) + 5
		fmt.Println(intn)
	}

}
