// 枚举iota
package main

import (
	"fmt"
)

/*
func main01() {
	const (
		a = iota //0  静止
		b = iota //1  移动
		c = iota //2  普通攻击
		d = iota //3
	)
	fmt.Println(a, b, c, d)
	//定义为状态
	value := a
	fmt.Println(value)
	value = b
	fmt.Println(value)
}
func main02() {
	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)
}*/
func main() {
	const (
		a    = iota
		b, c = iota, iota
		d, e = iota, iota
	)
	fmt.Println(a, b, c, d, e)
}
