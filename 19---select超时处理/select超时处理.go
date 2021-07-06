package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println(num)
			case <-time.After(3 * time.Second): //定时到达后，要处理的内容
				quit <- true
				goto lable
			}
		}
	lable:
		fmt.Println("break to lable---")
	}()
	for i := 0; i < 3; i++ {
		ch <- i
		time.Sleep(time.Second)

	}
	<-quit //主go程阻塞等待，子go程通知退出
	fmt.Println("finish")
}
