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

func (s *student) SayHii() {
	fmt.Printf("大家好，我是%s,我是%s生,我今年%d岁了,我的成绩是%d分\n", s.name, s.sex, s.age, s.score)
}
func (t *teacher) SayHii() {
	fmt.Printf("大家好，我是%s,我是%s生,我今年%d岁了,我教的科目是%s", t.name, t.sex, t.age, t.subject)
}

//接口的实现
type Person interface {
	SayHii()
}

//多态实现
//多态是将接口类型作为函数参数 多态实现了接口的统一处理
func SayHii(p Person) {
	p.SayHii()
}
func main() {
	var p Person
	p = &student{person{"大象", 19, "男"}, 65}
	SayHii(p)
	p = &teacher{person{"蚂蚁", 20, "男"}, "英语"}
	SayHii(p)
}
