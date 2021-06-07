// 水仙花
package main

import (
	"fmt"
)

func main01() {
	for i := 100; i < 1000; i++ {
		a := i / 100
		b := i % 100 / 10
		c := i % 10
		if a*a*a+b*b*b+c*c*c == i {
			fmt.Println(i)
		}
	}
}
func main() {
	for i := 1; i < 100; i++ {
		if i%7 == 0 || i%10 == 7 || i/10 == 7 {
			fmt.Println(i)
		}
	}
}
