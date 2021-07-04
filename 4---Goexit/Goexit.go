package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccccc")
	runtime.Goexit() //退出当前Go程
	fmt.Println("ddddddd")

}
func main() {
	go func() {
		defer fmt.Println("aaaaa")
		go test()
		fmt.Println("bbbbbb")
	}()
	for {

	}
}
