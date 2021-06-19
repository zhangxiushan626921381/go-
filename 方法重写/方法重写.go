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

//方法重写  在一个对象中不能出现相同的方法名 方法的接收者 带*和不带*表示一个相同的对象
func (p person) PrintInfo() {
	fmt.Printf("大家好我是%s,我今年%d岁,我是%s生", p.name, p.age, p.sex)
}
func (s student) PrintInfo() {
	fmt.Printf("大家好我是%s,我今年%d岁,我是%s生,我的成绩是%d", s.name, s.age, s.sex, s.score)
}
func main() {
	var s student = student{person{"大象", 16, "男"}, 56}
	//调用父类方法
	s.person.PrintInfo()
	//调用子类方法
	s.PrintInfo()
}
