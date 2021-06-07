// 输出格式 project main.go
package main

import (
	"fmt"
)

func main0401() {
	//输出格式
	//Println  打印数据时自动换行
	//Println  打印数据不换行
	fmt.Println("Hello World!")
	fmt.Println(10)
	fmt.Println(3.14)
	a := 10333333
	b := 3.145455555
	//%d 占位符 表示输出一个整形数据
	//%f 占位符 表示输出一个浮点型数据
	//\n 表示转义字符  相当于换行符
	//%f 占位符 默认保留6位小数
	//%3f  表示小数点后面保留三位 会对第四位进行四舍五入
	fmt.Printf("%d\n", a)
	fmt.Printf("%.3f\n", b)
}
func main() {
	//bool布尔 	string字符串 byte字符
	//声明bool类型变量 默认值为 false true
	//%t 占位符 表示输出一个布尔类型数据

	var a bool
	a = true
	fmt.Println(a)
	fmt.Printf("%t\n", a)
	var b string
	b = "性感荷官在线发牌"
	fmt.Println(b)
	//%s 占位符 表示输出一个字符串类型数据
	fmt.Printf("%s\n", b)
	var c byte
	c = 'a'
	//字符型变量对应的ASCII码值
	fmt.Println(c)
	//%c 占位符 表示输出一个字符类型数据
	fmt.Printf("%c", c)

}
