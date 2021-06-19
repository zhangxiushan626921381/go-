package main

import "fmt"

type p struct { //父结构体
	id   int
	name string
	age  int
	sex  string
}
type s struct { //子结构体
	m p //结构体变量  结构体类型
	//p			//继承父结构体
	score int
}

func (std *s) SayHello() { //父结构体的方法
	fmt.Println("大家好我是", std.m.name)
}

func main() {
	var stu s
	var std s = s{p{2, "蚂蚁", 18, "男"}, 18}
	//子类结构体继承父类结构体  允许使用父类结构体成员 允许使用父类的方法
	//父类不能继承子类信息
	std.SayHello()
	stu.score = 26
	stu.m.id = 1
	stu.m.name = "大象"
	stu.m.age = 16
	stu.m.sex = "男"
	fmt.Println(stu)

}
