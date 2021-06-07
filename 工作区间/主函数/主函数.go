package main

//如果调用不同级别目录需要导入包
import "工作区间/user"

//全局变量
var num int = 123

//在同级目录 包名要相同
func main() {
	add(10, 30)
	user.Login()
	user.Adduser()
	user.Delete()
}
