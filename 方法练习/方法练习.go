package main

import "fmt"

type student struct {
	name   string
	age    int
	sex    string
	cscore int
	mscore int
	escore int
}

//方法也是全局的，允许程序在所有文件访问
//为对象赋值
func (s *student) add(name string, age int, sex string, cscore int, mscore int, escore int) {
	s.name = name
	s.age = age
	s.sex = sex
	s.mscore = mscore
	s.escore = escore
	s.cscore = cscore
}

//方法1 打招呼
func (s *student) SayHello() {
	fmt.Printf("大家好,我叫%s，我今年%d岁了，我是%s生\n", s.name, s.age, s.sex)
}

//方法名可以和函数重名
func SayHello() {
	fmt.Println("hello world")
}

//方法2 打印成绩
func (s *student) PrintScore() {
	sum := s.cscore + s.escore + s.mscore
	fmt.Printf("我的总成绩为%d,我的平均成绩为%d\n", sum, sum/3)
}
func main() {
	var s student = student{"蜗牛", 18, "男", 18, 19, 23}
	s.SayHello()
	s.PrintScore()
	//初始化对象信息
	s.add("大象", 15, "男", 23, 56, 89)
	s.SayHello()
	s.PrintScore()
	SayHello()

}
