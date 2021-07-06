package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var RWMutex sync.RWMutex

func read(in <-chan int, idx int) {
	//RWMutex.RLock() //以读模式加锁
	for {
		num := <-in
		fmt.Printf("第%d个读go程读出%d\n", idx, num)
	}
	//RWMutex.RUnlock() //以读模式解锁
}
func write(out chan<- int, idx int) {
	RWMutex.Lock() //以写模式加锁
	for {
		//生成随机数
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("第%d个写go程写入%d\n", idx, num)
		time.Sleep(time.Millisecond)
	}
	RWMutex.Unlock() //以写模式解锁
}
func main() {
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())
	/*quit:=make(chan bool)//用于关闭主go程的channel*/
	ch := make(chan int) //用于数据传递的channel
	for i := 0; i < 5; i++ {
		go read(ch, i)
	}
	for i := 0; i < 5; i++ {
		go write(ch, i)
	}
	for {

	}
}
