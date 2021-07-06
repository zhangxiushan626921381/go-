package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ch = make(chan int)

func read(in <-chan int, idx int) {
	for {
		in := <-ch //从channel中读数据
		fmt.Printf("第%d个读go程读出%d\n", idx, in)
		time.Sleep(time.Second)
	}

}
func write(out chan<- int, idx int) {
	for {
		//生成随机数
		num := rand.Intn(1000)
		out <- num //写入 channel
		fmt.Printf("第%d个写go程写入%d\n", idx, num)
		time.Sleep(time.Millisecond * 300)

	}
}
func main() {
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		go read(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go write(ch, i+1)
	}
	for {

	}
}
