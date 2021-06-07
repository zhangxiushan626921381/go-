// 闭包
package main

import (
	"fmt"
)

//函数在调用结束会从内存中销毁
func test1(a int) {
	a++
	fmt.Println(a)
}

func main01() {
	a := 0
	for i := 0; i < 10; i++ {
		test1(a)
	}
}

//可以通过匿名函数和闭包实现函数在栈区的本地化操作
func test2() func() int {
	var a int
	return func() int {
		a++
		return a
	}
}
func main() {
	//将test2函数类型赋值给f
	//f := test2   返回值类型是func func ()int
	//函数调用将test2的返回值给f
	f := test2() //返回值类型是func()int
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	fmt.Printf("%T", f)
}
