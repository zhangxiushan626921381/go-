package main

import "fmt"

func main() {
	ch := make(chan string) //无缓冲channel
	//len(ch)求取的是channel中剩余未读取数据个数
	//cap(ch)求取的是通道的容量
	fmt.Println("len(ch)", len(ch), "cap(ch)", cap(ch))
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("i:", i)
		}
		//通知主Go程打印完毕
		ch <- "子Go程打印完毕"
	}()
	str := <-ch
	fmt.Println(str)
}
