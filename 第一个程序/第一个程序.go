// 第一个程序 project main.go
// 导入主函数的包
package main

// 系统会导入所需的包   format 标准输入输出包
import (
	"fmt"
)

// 注释不参与程序编译，可以帮助理解程序
// main叫主函数 是程序的主入口 程序有且只有一个主函数
//行注释
/*
块注释
可以注释多行
*/
//快速注释的快捷键 ctrl+/    选中所需要的注释的内容

func main() {
	//在屏幕打印helloworld
	fmt.Println("helloworld")
	fmt.Println("zhangxiushan")

}
