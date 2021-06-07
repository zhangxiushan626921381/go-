// 函数的定义和使用
package main

import (
	"fmt"
)

func add(a int, b int) {
	var sum int
	sum = a + b
	fmt.Println(sum)
}

func main() {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	add(a, b)

}
