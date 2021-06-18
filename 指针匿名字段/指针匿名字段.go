package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}
type student struct {
	*person //指针匿名字段
	score   int
	id      int
}

func main() {
	var stu student
	stu.person = new(person)
	var per = person{"杨过", 16, "男"}
	//将结构体变量赋值给结构体指针类型
	stu.person = &per
	stu.id = 106
	stu.score = 100
	fmt.Println(stu)
	fmt.Println(stu.name)
	var std student = student{&person{"ss", 18, "男"}, 103, 56}
	fmt.Println(std)
	fmt.Println(std.name)
}
