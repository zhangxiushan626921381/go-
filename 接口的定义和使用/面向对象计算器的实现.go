package main

import "fmt"

type opt struct {
	num1 int
	num2 int
}
type SubOpt struct {
	opt
}
type AddOpt struct {
	opt
}
type C struct {
	opt
}

func (sub *SubOpt) operate() int {
	return sub.num1 - sub.num2
}
func (add *AddOpt) operate() int {
	return add.num1 + add.num2
}
func (cf *C) operate() int {
	return cf.num1 / cf.num2
}

//定义接口
type Opter interface {
	operate() int
}

func main() {
	/*var add AddOpt=AddOpt{opt{10,20}}
		var sub SubOpt=SubOpt{opt{10,20}}
	 	add.operate()
		sub.operate()
	*/
	var o Opter
	o = &AddOpt{opt{10, 20}}
	value := o.operate()
	fmt.Println(value)
	o = &SubOpt{opt{10, 20}}
	value = o.operate()
	fmt.Println(value)
	o = &C{opt{20, 20}}
	value = o.operate()
	fmt.Println(value)
}
