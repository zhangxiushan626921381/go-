// 占位符的使用
package main

import (
	"fmt"
)

func main01() {
	var a int = 10
	fmt.Printf("%d\n", a)
	var b float64 = 10
	fmt.Printf("%f\n", b)
	var c bool = true
	fmt.Printf("%t\n", c)
	var d byte = 'b'
	fmt.Printf("%c\n", d)
	var e string = "hello world"
	fmt.Printf("%s\n", e)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%T\n", a)

}
func main() {
	//计算机能够识别的进制  二进制  八进制  十进制 十六进制
	a := 123   //十进制数据
	b := 0123  //八进制数据  以0开头
	c := 0x123 //十六进制数据  以0x开头
	fmt.Println(a, b, c)
	fmt.Printf("a的二进制%b", a)
}
