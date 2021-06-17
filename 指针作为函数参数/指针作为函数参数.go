package main

import "fmt"

func swap(a int, b int) {
	a, b = b, a
}
func main01() {
	a := 10
	b := 20
	//不能交换 a，b的值 原因是值传递
	swap(a, b)
	fmt.Println(a, b)
}
func swap1(a *int, b *int) {
	*a, *b = *b, *a
}
func main() {
	a := 10
	b := 20
	//指针作为函数参数 是地址传递
	swap1(&a, &b)
	fmt.Println(a, b)
}
