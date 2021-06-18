package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main01() {
	result := add(5, 4)
	fmt.Println(result)
}
