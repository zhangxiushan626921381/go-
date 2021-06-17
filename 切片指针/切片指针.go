package main

import "fmt"

func main() {
	var slice []int = []int{1, 2, 3, 4, 5}
	var p *[]int //二级指针
	p = &slice
	//切片名本身就是一个地址
	slice[2] = 200
	(*p)[3] = 300
	//增加元素
	*p = append(*p, 6, 7)
	fmt.Printf("%p\n", slice)
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", *p)
	fmt.Println(slice)
	fmt.Println(*p)
}
