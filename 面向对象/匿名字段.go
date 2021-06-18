package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}
type student struct {
	//通过匿名字段实现继承操作
	person //结构体名称作为结构体成员
	id     int
	score  int
}

func main01() {
	var stu student
	stu.name = "张三"
	stu.score = 100
	stu.age = 18
	stu.id = 101
	stu.sex = "男"
	fmt.Println(stu)
	var std student = student{person{"张四", 15, "男"}, 103, 520}
	fmt.Println(std)
}
