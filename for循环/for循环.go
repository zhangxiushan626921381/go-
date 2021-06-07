// for循环
package main

import (
	"fmt"
)

func main01() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
func main02() {
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum = sum + i
		}
	}
	fmt.Println(sum)
}
func main03() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
func main04() {
	i := 0
	for ; ; i++ {
		if i > 10 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println(i)
}
func main() {
	for {
		fmt.Println("hello")
	}
}
