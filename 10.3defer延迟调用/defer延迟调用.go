package main

import "fmt"

func main01() {
	fmt.Println("大象")
	defer fmt.Println("蚂蚁")
	//defer函数调用 defer调用的函数并没有直接使用  而是先加载到栈区内存中，在函数结束时进行调用 栈是先进后出原则
	//先打印青蛙 后打印蚂蚁
	defer fmt.Println("青蛙 ")
	fmt.Println("鸭子")
}
func test() {
	fmt.Println("hello world")
}

var value int

func test1(a int, b int) int {
	value = a + b
	return value
}
func test2(value int) {
	fmt.Println(value)
}
func main02() {
	test()
	//函数中有返回值不能使用defer调用
	defer test1(10, 20)
	test2(value)
}

func main() {
	a := 10
	b := 20
	//定义函数类型变量
	defer func() {
		fmt.Println(a)
		fmt.Println(b)
	}()
	a = 100
	b = 200
	fmt.Println(a, b)
}
