// 常量的定义和使用
package main

import (
	"fmt"
)

func main() {
	//常量的定义和使用
	const a int = 10
	const B int = 30

	var b int = 10
	var c int = 20
	fmt.Println(&b, &c)
	d := B + c
	//	a = 20
	fmt.Println(a, d)
	fmt.Println("Hello World!")
}
