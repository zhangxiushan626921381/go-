// 循环嵌套
package main

import (
	"fmt"
)

func main01() {
	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
		}

	}
}
func main0401() {

	//嵌套循环中  执行次数为外层*内层
	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
		}

	}
}
func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)

		}
		fmt.Println()
	}
}
