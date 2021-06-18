package main

import "fmt"

//定义函数类型
//为存在的数据类型起别名
type INT int

//方法
//func （方法接收者）方法名（参数列表）返回值类型
func (a INT) add(b INT) INT {
	return a + b
}
func main02() {
	//根据数据类型绑定方法
	var a INT = 30
	value := a.add(20)
	fmt.Println(value)
}
