package main

import (
	"fmt"
	"time"
)

//定义全局channel，用来完成数据同步操作
var channel = make(chan int)

//定义一台打印机
func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch) //屏幕:stdout
		time.Sleep(2000 * time.Millisecond)
	}
}

//定义两个人使用打印机
//先执行person1没有写入8通道就一直是阻塞的，因为没法读，读和写都会导致通道的阻塞
//所以就不会执行person2  所以就会先打印完hello然后写入8
func person1() { //person1先执行
	printer("hello")
	channel <- 8
}
func person2() { //person2后执行
	<-channel //<-channel1
	printer("你好")

}
func main() {
	//go程的特性：主结束 子结束
	go person1()
	go person2()
	for {

	}
}
