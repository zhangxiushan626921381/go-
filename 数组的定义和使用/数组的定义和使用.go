package main

import (
	"fmt"
)

var acc [10]int = [10]int{1, 5, 8, 9, 6, 2, 3, 5, 8}

//打印数组信息 len(string)获取字符串的有效长度
//len(...int)获取不定参函数参数个数
//len([10]int)计算数组的元素个数
//rang遍历一个集合的信息
func main01() {
	/*for i:=0;i<len(acc);i++{
		fmt.Println(acc[i])
	}*/
	//i 表示下标 v 表示具体的值
	//不想接收，用_来表示
	for i, data := range acc {
		fmt.Println(i, data)

	}
}
func main02() {
	//数组在定义时  可以初始化部分元素的值
	//使用自动类型推导创建数组
	ayy := [10]int{6, 5, 9, 59}

	for i, data := range ayy {
		fmt.Println(i, data)
	}
	//打印数组的类型
	fmt.Printf("%T", ayy)
}
func main() {
	//5 在定义时叫元素个数
	//数组在定义好后就不能再添加数组了 元素个数固定 不能添加
	//数组是一个常量 不允许赋值操作
	atu := [5]int{5, 9, 8, 3, 1}
	//在数组使用时叫下标 （0~5-1)
	fmt.Println(atu)

}
