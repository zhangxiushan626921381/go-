// 函数返回多个值
package main

import (
	"fmt"
)

func max(a int, b int) {
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}

}
func main() {
	max(3, 5)
}
