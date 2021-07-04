package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	name string
	age  int
	id   int
}
type student struct {
	person
	score []int
}

func test(man person) {
	man.name = "mike"
	man.age = 15
	man.id = 001
}
func main02() {
	var p1 *person = &person{"mike", 14, 10}
	fmt.Println(*p1)
	var temp person = person{"tom", 46, 56}
	var p2 *person
	p2 = new(person)
	p2 = &temp
	fmt.Println(*p2)
	var p3 person
	p3 = person{"d", 56, 22}
	fmt.Println(p3)

}
func main01() {
	var temp person
	v := unsafe.Sizeof(temp)
	test(temp) //值传递  将实参的值拷贝一份给形参
	fmt.Println(temp, v)
}

type person1 struct {
	name        string
	age         int
	flg         bool
	interesting []string
}

func initFun(p *person1) {
	p.age = 59
	p.name = "tom"
	p.flg = true
	p.interesting = []string{"唱歌", "跳舞"}
}
func main() {
	var person person1
	initFun(&person)
	fmt.Println(person)
}
