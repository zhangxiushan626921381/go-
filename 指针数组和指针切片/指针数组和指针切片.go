package main

import "fmt"

func main01() {
	// var arr [3]int    //int数据
	var arr [3]*int //*int数据
	a := 10
	b := 20
	c := 30
	arr[0] = &a
	arr[1] = &b
	arr[2] = &c
	//通过指针数组改变变量的值
	*arr[2] = 60
	fmt.Println(c)
	for i := 0; i < len(arr); i++ {
		fmt.Println(*arr[i])
	}
}
func main() {
	//指针切片
	var slice []*int
	a := 10
	b := 20
	c := 30
	d := 50
	//通过地址添加
	slice = append(slice, &a, &b, &c)
	slice = append(slice, &d)
	fmt.Println(slice)
	for _, v := range slice {
		fmt.Println(*v)
	}
}
