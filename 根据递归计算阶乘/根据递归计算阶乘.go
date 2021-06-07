// 根据递归计算阶乘
package main

import (
	"fmt"
)

var sum int = 1

func test(n int) {
	if n == 5 {
		return
	}
	n++
	test(n)
	sum *= n

}

func main() {
	test(3)
	fmt.Println(sum)

}
