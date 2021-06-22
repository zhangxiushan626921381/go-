package main

import "fmt"

//接口
type Humaner interface { //子集
	SayHi()
}

//接口
type Person interface { //超集
	Humaner //一组子集的集合
	Sing(string)
}
type student struct {
	name string
	age  int
	sex  string
}

func (s *student) SayHi() {
	fmt.Printf("大家好，我是%s,我今年%d岁了，我是%s生\n", s.name, s.age, s.sex)
}
func (s *student) Sing(name string) {
	fmt.Printf("大家好，我是%s,我给大家唱一首%s\n", s.name, s.name)
}
func main() {
	var h Humaner
	h = &student{"韩红", 45, "女"}
	h.SayHi()
	var p Person //超集
	p = &student{"大象", 46, "男"}
	p.SayHi()
	p.Sing("我爱祖国")
	h = p //将超集转成子集  反过来是不允许的
	h.SayHi()
}
