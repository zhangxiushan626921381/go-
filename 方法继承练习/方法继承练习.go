package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}
type jizhe struct {
	person
	aihao string
	job   string
}
type chengxuyuan struct {
	person
	jobtime int
}

func (p person) SayHello() {
	fmt.Printf("我是%s,我今年%d岁,我是%s生,", p.name, p.age, p.sex)
}
func (jz *jizhe) SayHello() {
	jz.person.SayHello()
	fmt.Printf("我的爱好是%s,我是一个%s狗仔\n", jz.aihao, jz.sex)
}
func (cxy *chengxuyuan) SayHello() {
	cxy.person.SayHello()
	fmt.Printf("我的工作年限是%d年", cxy.jobtime)
}
func main() {
	var jz jizhe = jizhe{person{"记者", 34, "男"}, "拍照", "记者"}
	jz.SayHello()
	var cxy chengxuyuan = chengxuyuan{person{"汤姆逊", 68, "男"}, 40}
	cxy.SayHello()
}
