// 多重赋值 project main.go
package main

import (
	"fmt"
)

func main0301() {
	//格式 变量1，变量2:=值1,值2
	a, b, c, d := 10, 20, 30, 40
	fmt.Println(a, b, c, d)
}
func main0302() {
	a, b := 10, 20
	var temp int
	//交换a b的值
	//第一种方法：
	temp = a
	a = b
	b = temp
	//第二种方法：
	a = a + b
	b = a - b
	a = a - b
	//第三种方法：
	a, b = b, a
	fmt.Println(a, b)
}

//匿名变量   接受这个数据 但是不做存储
func main0303() {
	_, b, c := 1, 2, 3
	//fmt.Println(_)
	fmt.Println(b)
	fmt.Println(c)
}
