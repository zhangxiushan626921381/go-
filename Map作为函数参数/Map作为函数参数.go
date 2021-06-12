package main

import "fmt"

func demo(m map[int]string) {
	//map在函数中添加数据是可以的 影响主调函数中实参的值
	m[106] = "沙师弟"
	m[105] = "猪八戒"
	fmt.Println(len(m))
	delete(m, 101)
	fmt.Println(len(m))
}
func main() {
	m := make(map[int]string, 1)
	m[101] = "孙悟空"
	fmt.Println(len(m))
	//map作为函数参数是引用传递（地址传递）
	demo(m)
	fmt.Println(m)
}

//m:=make(map[int][]int) //key是int类型 value是切片类型

//m:=make(map[int]struct) //key是int类型 value是结构体类型
