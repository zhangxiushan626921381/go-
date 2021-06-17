package main

import "fmt"

func main() {
	var arr [5]int = [5]int{1, 5, 6, 8, 9}
	//定义指针指向数组
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])
	fmt.Println(arr)
	var p *[5]int
	//将指针变量和数组建立关系
	p = &arr
	//通过指针找数组值
	fmt.Println((*p)[0])
	//直接使用指针[下标]操作数组元素
	fmt.Println(p[2])
	//可以通过指针间接操作数组
	*p = [5]int{5, 6, 7, 8, 8}
	fmt.Println(arr)
	fmt.Printf("%p\n", p)
}
