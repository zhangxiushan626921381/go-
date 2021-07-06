package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int) //用来进行数据通信的channel
	ch1 := make(chan int)
	quit := make(chan bool) //用来判断是否退出的channel
	go func() {             //子go程写数据
		for i := 0; i < 3; i++ {
			ch <- i
			ch1 <- i
			time.Sleep(time.Second)
		}
		close(ch)
		close(ch1)
		quit <- true //通知主Go程退出
		runtime.Goexit()
	}()
	for { //主go程读数据
		select {
		case num := <-ch:
			fmt.Println("我从ch读到数据：", num)
		case num := <-ch1:
			fmt.Println("我从ch1 读到数据：", num)
		case <-quit:
			return
		}
		fmt.Println("------------------------")
	}
}
