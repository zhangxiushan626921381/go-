// switch分支语句
package main

import (
	"fmt"
)

func main() {
	var w int
	fmt.Scan(&w)
	switch w {
	case 1:
		fallthrough
	case 3:
		fallthrough
	case 5:
		fallthrough
	case 12:
		fmt.Println("31天")
	default:
		fmt.Println("输入错误")
	}
}
