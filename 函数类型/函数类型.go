// 函数类型
package main

import (
	"fmt"
)

/*
func test1() {
	fmt.Println("askdkasj")
}
func test2(a int, b int) {
	fmt.Println(a + b)
}
func test3(a int, b int) int {
	return a * b
}

type TYPE1 func()
type TYPE2 func(int, int)
type TYPE3 func(int, int) int

func main() {
	var f TYPE1

	f = test1
	f()
	var f1 TYPE2
	var f2 TYPE3
	f2 = test3
	v := f2(3, 5)
	fmt.Println(v)

	f1 = test2
	f1(10, 30)
}*/
func test1(a int, b int) {
	fmt.Println(a + b)
}
func main() {
	f := test1
	f(10, 30)
}
