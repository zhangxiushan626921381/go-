package main

import (
	"fmt"
	"runtime"
)

func main01() {

	n := runtime.GOMAXPROCS(1) //将cpu设置为单核
	fmt.Println("n:", 1)       //n返回的是上一次设置的cpu核数
	n = runtime.GOMAXPROCS(2)  //将cpu设置为双核
	fmt.Println("n:", n)
	for {
		go fmt.Println(0) //子Go程   直接不断的打印0
		fmt.Println(1)    //主Go程
	}
}
func main() {
	k := runtime.GOROOT() //返回GO的根目录
	fmt.Println(k)
}
