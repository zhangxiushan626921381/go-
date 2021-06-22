package main

import "fmt"

func demo(i int) {
	defer func() {
		//recover可以从panic异常中重新获取控制权
		//recover()  返回值为接口类型
		//fmt.Println(recover())
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("11")
	var p *int
	*p = 123 //err   recover作用范围
	fmt.Println("22")
	var arr [10]int
	//recover使用一定要在错误出现之前进行拦截
	//通过匿名函数和recover进行错误的拦截

	//如果传递超出数组下标值 导致数组下标越界
	arr[i] = 100 //panic 系统处理 导致程序崩溃
	fmt.Println(arr[i])
}
func demo1() {
	fmt.Println("hello world")
}
func main() {
	demo(10)
	demo1()
}
