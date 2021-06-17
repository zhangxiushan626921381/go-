package main

import "fmt"

func main() {
	var p *int //int类型指针只能开辟int类型
	fmt.Println(p)
	//为指针变量创建一块内存空间
	//在堆区创建空间
	p = new(int) // new 创建好的空间值为数据类型的默认值
	//查看创建的空间
	//在Go语言中只需要开辟空间new（数据类型）不需要管理空间的释放
	//打印P指向空间的值
	*p = 123
	fmt.Println(*p)
	//打印P的值
	fmt.Println(p)
}
