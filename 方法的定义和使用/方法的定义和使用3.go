package main

import "fmt"

type student struct {
	name string
	age  int
	sex  string
}

//方法 值传递
func (s student) Printinfo() {
	fmt.Println(s.name)
	fmt.Println(s.age)
	fmt.Println(s.sex)
}

//方法 地址传递
func (s *student) edit(name string, age int, sex string) {
	//(*s).name=name
	//结构体指针可以直接调用结构体的成员
	s.name = name
	s.age = age
	s.sex = sex
}
func main() {
	//var s student=student{"鸭子",15,"ss"}
	//对象.方法
	//s.Printinfo()
	var s1 student = student{"企鹅", 16, "sss"}
	s1.Printinfo()
	//变量可以直接使用结构体指针对应的方法
	s1.edit("海豹", 18, "ssd")
	s1.Printinfo()
}
