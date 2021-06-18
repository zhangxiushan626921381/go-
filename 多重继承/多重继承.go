package main

import "fmt"

type A struct {
	name string
	id   int
}
type B struct {
	A   //B继承A
	sex string
	age int
}

//注意：结构体不能嵌套本结构体
//结构体可以嵌套本结构体指针类型  链表
type C struct {
	B     //C继承B
	score int
	add   string
}

func main1() {
	var stu C
	stu.name = "李四"
	stu.id = 101
	stu.age = 20
	stu.score = 10
	stu.sex = "男"
	fmt.Println(stu)
}
