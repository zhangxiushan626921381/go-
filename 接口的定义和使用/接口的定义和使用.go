package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}
type student struct {
	person
	score int
}
type teacher struct {
	person
	subject string
}

func (s *student) SayHi() {
	fmt.Printf("大家好，我是%s,我是%s生,我今年%d岁了,我的成绩是%d分\n", s.name, s.sex, s.age, s.score)
}
func (t *teacher) SayHi() {
	fmt.Printf("大家好，我是%s,我是%s生,我今年%d岁了,我教的科目是%s", t.name, t.sex, t.age, t.subject)
}

//接口定义
//接口定义了规则  方法实现了规则
//接口是虚的  方法是实的
//格式：type 接口名 interface{方法列表}
//方法名(参数列表)(返回值列表)
type Humaner interface {
	//方法/函数的声明没有具体实现 具体的实现要根据对象不同  实现的方式也不同
	//接口中的方法必须有具体的实现
	SayHi()
}

func main() {
	var s student = student{person{"大象", 19, "男"}, 65}
	var t teacher = teacher{person{"蚂蚁", 20, "男"}, "英语"}
	//	定义接口类型
	//接口做了统一的处理 先实现接口再根据接口实现对应的方法
	//在需求改变时，不需要改变接口 只需要修改方法
	var h Humaner
	h = &s
	h.SayHi()
	h = &t
	h.SayHi()
	//t.SayHi()
	//s.SayHi()
}
