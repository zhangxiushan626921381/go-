// 01字符和字符串的区别
package main

import (
	"fmt"
)

/*
func main() {
	var b string = "ajjjj"

	fmt.Printf("%s", b)
}*/
func main() {
	//var a string = "hello world"
	//计算字符串的个数
	var b string = "传智播客"
	//在go语言中汉子是3个字符，为了和linux同一处理
	//num := len(a)
	num := len(b)
	fmt.Println(num)
	fmt.Println(num)
}
