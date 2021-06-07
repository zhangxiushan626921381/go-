// 匿名函数
package main

import (
	"fmt"
)

func main01() {
	a := 10
	b := 20
	//匿名内部函数
	func(a int, b int) {
		fmt.Println(a + b)
	}(a, b)

}
func main02() {
	a := 10
	b := 20
	//f是函数类型对应的变量
	f := func(a int, b int) {
		fmt.Println(a + b)
	}
	f(a, b)
	f(20, 30)
}
func main() {
	a := 10
	b := 20
	//匿名函数  带返回值的匿名函数
	f := func(a int, b int) int {
		return a + b
	}
	//函数调用	f 函数类型 v 接受返回值  类型为int
	v := f(a, b)
	fmt.Println(v)
}
