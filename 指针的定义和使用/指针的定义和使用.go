package main

import "fmt"

func main01() {
	//定义整型指针变量 指向a的地址
	//指针类型定义
	//var 指针 *数据类型  一级指针
	var p *int
	a := 10
	//将a的地址赋值给指针变量p
	p = &a
	//通过指针变量 间接访问变量对应的内存空间
	fmt.Println(*p)
	//可以通过指针间接改变变量
	*p = 100
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(&a)
}
func main() {
	/*a:=10
	p:=&a
	fmt.Println(p)
	*/
	//声明了一个指针  默认值是nil（空指针 值为0）指向了内存地址编号为0的空间
	//0-255为系统占用  不允许用户进行读写操作
	var p *int //空指针
	//p=0xc00000a0a0 //野指针  指针变量	指向了一个未知的空间 访问野指针和空指针对应的内存空间都会报错
	b := 123
	p = &b
	fmt.Println(p)
}
