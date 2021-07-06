package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool) //判断一个是否终止的channel
	fmt.Println(time.Now())
	myticker := time.NewTicker(time.Second * 1) //定义一个定时器
	i := 0
	go func() {
		for {
			nowtime := <-myticker.C
			fmt.Println(nowtime)
			i++
			if i == 8 {
				quit <- true //解除主go程阻塞
				break
			}
		}

	}()
	<-quit //子go程循环获取<-myTicker.C期间，一直阻塞
}
