// 逻辑运算符
package main

import (
	"fmt"
)

func main01() {
	var a bool = false
	a = true
	b := 20
	c := 20
	fmt.Println(!(b == c))
	fmt.Println(!a)
}
func main02() {
	a := 10
	b := 20

	fmt.Println(a < b && false)

}
func main03() {
	a := 10
	b := 20
	c := 30
	d := 20
	fmt.Println(a < b || c < d)
}
func main() {
	//取地址运算符 &
	var a int = 10
	fmt.Println(&a)
	//取值运算符
	b := 20
	p := (&b)

	fmt.Println(*p)
}
