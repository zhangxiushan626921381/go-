package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var RWMutex sync.RWMutex
var value int //定义全局变量 模拟共享数据
func read(idx int) {
	for {
		RWMutex.RLock() //以读模式加锁
		num := value
		fmt.Printf("第%d个读go程读出%d\n", idx, num)
		RWMutex.RUnlock() //以读模式解锁
		time.Sleep(time.Second)
	}

}
func write(idx int) {
	for {
		//生成随机数
		num := rand.Intn(1000)
		RWMutex.Lock() //以写模式加锁
		value = num
		fmt.Printf("第%d个写go程写入%d\n", idx, num)
		time.Sleep(time.Millisecond * 300)
		RWMutex.Unlock() //以写模式解锁
	}
}
func main() {
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		go read(i + 1)
	}
	for i := 0; i < 5; i++ {
		go write(i + 1)
	}
	for {

	}
}
