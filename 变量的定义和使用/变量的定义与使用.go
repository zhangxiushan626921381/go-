// 变量的定义和使用 project main.go
package main

import (
	"fmt"
)


func main01() {
	//变量的定义和说明
	//声明和定义
	//var 变量名  数据类型
	//布尔类型 整型 浮点型 字符型 字符串类型
	//变量在定义时 存储不同的数据  需要定义不同的数据类型
	var a int = 10
	//声明，如果没有初始值，默认值为0

	a = a + 25

	fmt.Println(a)

}
func main02() {
    计算圆的面积和周长 	 //自动推导类型
	
 	PI :=float64
	// s = PI * r * r
	s := PI * r * r
	l = PI * 2 * r	
	wty := "你愁啥	"
 	fmt.Println("面积：", s)
 	fmt.Println("周长：", l)
 	fmt.Println(wty)


func main03() {
    //自动推到类型 变量名：=值   会根据变量选择类型
	//买两斤黄瓜   价格为5元
	w := 2.0 //float64
	// p := 5  //int
	//在go语言中不同的数据类型不能计算 可以通过类型转换解决
	var p float64 = 5
	fmt.Println(p * w)

}
func main() {
	a := false  //bool
	b := 10     //int
	c := 3.14   //float64
	d := 'a'    //byte字节类型
	e := "瞅你咋地" //string 字符串类型
	fmt.Println(a, b, c, d, e)

}
