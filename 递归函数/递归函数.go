// 递归函数
package main

import (
	"fmt"
)

//在函数定义时，调用本身叫做递归函数
//死递归
func test3(a int) {
	//在递归函数中要有一个出口 用return来表示
	if a == 0 {
		return
	}
	a--
	fmt.Println(a)
	test3(a)

}
func main() {
	test3(10)

}
