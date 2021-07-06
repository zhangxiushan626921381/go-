package main

import (
	"fmt"
	"time"
)

func main01() {
	fmt.Println("当前时间", time.Now())
	//创建定时器
	myTimer := time.NewTimer(time.Second * 2)
	nowTime := <-myTimer.C //chan类型
	fmt.Println("两秒之后", nowTime)
}

//三种定时方法
func main() {
	//1.sleep
	time.Sleep(time.Second)
	//2.Timer.C
	myTimer := time.NewTimer(time.Second * 2) //创建定时器，指定定时时长
	nowTime := <-myTimer.C                    //chan类型  定时满，系统自动写入时间
	fmt.Println("两秒之后", nowTime)
	//3.time.After
	nowtim := <-time.After(time.Second * 2)
	fmt.Println("当前时间", nowtim)
}
