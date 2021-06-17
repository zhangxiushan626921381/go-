package main

import "fmt"

type student struct {
	id   int
	name string
	age  int
	sex  string
}

func main01() {
	//定义结构体变量
	//var 结构体名 结构体数据类型
	var stu student = student{101, "哆啦A梦", 12, "男"}

	//结构体变量.成员=值
	//定义结构体指针指向变量的地址
	//var p *student
	//结构体指针指向结构体变量的地址
	//p=&stu
	fmt.Printf("%p\n", &stu)    //结构体类型
	fmt.Printf("%p\n", &stu.id) //int类型
}
func main() {
	var stu student = student{101, "哆啦A梦", 12, "男"}
	var p *student
	p = &stu
	//通过结构体指针间接操作结构体成员
	(*p).id = 102
	//通过指针可以直接操作结构体成员
	p.name = "大熊"
	fmt.Println(stu)
}
