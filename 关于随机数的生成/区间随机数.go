package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a int
	rand.Seed(time.Now().UnixNano())
	a = rand.Intn(10) + 10
	b := rand.Intn(2) + 10
	fmt.Println(a, b)
}
