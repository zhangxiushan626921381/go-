package main

import "fmt"

type person1 struct { //父类
	id   int
	name string
	age  int
}
type classmate struct { //子类
	person1
	name  string
	score int
}

func main() {
	var stu classmate
	//采用就近原则 使用子类信息
	stu.name = "张三丰"
	//结构体变量.匿名字段.结构体成员
	stu.person1.name = "张四丰"
	stu.id = 101
	stu.score = 50
	stu.age = 12
	fmt.Println(stu)
	var std classmate = classmate{person1{102, "嘎嘎", 16}, "嘎嘎嘎", 20}
	fmt.Println(std)
}
