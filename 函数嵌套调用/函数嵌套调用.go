// 函数嵌套调用
package main

import (
	"fmt"
)

func test1(arr ...int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}
}
func test2(b ...int) {
	//test1(b[0:]...)
	test1(b[1:3]...)
}
func main() {
	test2(1, 5, 8, 6, 9)
}
