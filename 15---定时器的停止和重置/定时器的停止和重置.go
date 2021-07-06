package main

import (
	"fmt"
	"time"
)

func main() {
	myTime := time.NewTimer(time.Second * 3) //创建一个定时器3
	myTime.Reset(10 * time.Second)           //重置定时时长为10
	go func() {
		<-myTime.C
		fmt.Println("子go程定时完毕")
	}()
	/*myTime.Stop()  //设置定时器停止*/
	for {

	}
}
