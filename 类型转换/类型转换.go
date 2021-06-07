// 类型转换
package main

import (
	"fmt"
)

func main01() {
	var a int = 10
	var b float64 = 3.14
	//将不同类型转成相同类型 进行计算的操作
	c := float64(a) * b
	fmt.Println(c)
}
func main() {

	fmt.Printf("%d周%d天", 46/7, 46%7)
	//编程实现107653是几天几时几分几秒
	fmt.Printf("%d天%d时%d分%d秒", 107653/86400%365, 107653/3600%24, 107653/60%60, 107653%60)
}
