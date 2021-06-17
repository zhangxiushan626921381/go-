package main

import "fmt"

func main() {
	a := 10
	p := &a
	m := &p //二级指针  二级指针存储一级指针的地址
	k := &m //三级指针  三级指针存储二级指针的地址
	//指针变量前加*表示具体的值
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", m)
	fmt.Printf("%p\n", *k)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", m)
	fmt.Printf("%T\n", k)
}
