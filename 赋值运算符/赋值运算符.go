// 赋值运算符
package main

import (
	"fmt"
)

func mains() {
	a := 10
	a = a + 5
	a -= 5
	a *= 5
	a %= 5
	a += 5
	fmt.Println(a)
}
func main() {
	a := 10
	a %= 5
	fmt.Println(a)
}
