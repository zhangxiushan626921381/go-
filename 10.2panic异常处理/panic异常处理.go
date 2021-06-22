package main

import "fmt"

func test1() {
	fmt.Println("hello world1")
}
func test2() {
	//可以在程序中直接调用Panic  调用之后程序会终止执行
	panic("hello world2")
	//print 和 println是对错误信息进行处理的  不建议在程序中打印数据进行使用
	//fmt.Println("hello world2")
}
func test3() {
	fmt.Println("hello world3")
}
func main() {
	test1()
	test2()
	test3()
}
