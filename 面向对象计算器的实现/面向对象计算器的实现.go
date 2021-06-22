package main

import "fmt"

type opt struct { //父类
	num1 int
	num2 int
}
type SubOpt struct { //减法子类
	opt
}
type AddOpt struct { //加法子类
	opt
}
type C struct { //除法子类
	opt
}

//减法方法实现
func (sub *SubOpt) operate() int {
	return sub.num1 - sub.num2
}

//加法方法实现
func (add *AddOpt) operate() int {
	return add.num1 + add.num2
}

//除法方法实现
func (cf *C) operate() int {
	return cf.num1 / cf.num2
}

//定义接口
type Opter interface {
	operate() int
}

//多态实现
func Fratory(o Opter) (value int) {
	value = o.operate()
	return value
}

//设计模式 对于面向对象 基于M（模型）V（视图）C（控制器）有26种
//工厂设计模式
//对象 空的
type OptFactory struct {
}

//num1 值1 num2 值2 op 运算符
func (of *OptFactory) OptCalc(num1 int, num2 int, op string) (value int) {
	//通过运算符进行分类计算
	//通过接口进行统一处理
	var opter Opter
	//根据不同的运算符
	switch op {
	case "+":
		//创建加法对象
		var add AddOpt = AddOpt{opt{num1, num2}}
		//value=add.operate()
		opter = &add
	case "-":
		//创建减法对象
		var sub SubOpt = SubOpt{opt{num1, num2}}
		opter = &sub
	case "/":
		var c C = C{opt{num1, num2}}
		opter = &c
	}
	//value=opter.operate() //通过多态实现接口的操作
	value = Fratory(opter)
	return
}

func main() {
	//1.基于继承和方法
	/*var add AddOpt=AddOpt{opt{10,20}}
		var sub SubOpt=SubOpt{opt{10,20}}
	 	add.operate()
		sub.operate()
	*/
	//2.基于继承方法和接口
	/*
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
	*/
	//3.基于继承 方法 接口 多态 设计模式
	var optf OptFactory
	value := optf.OptCalc(10, 20, "+")
	fmt.Println(value)
}
