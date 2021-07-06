package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond // 定义全局条件变量

func producer08(out chan<- int, idx int) {
	for {
		// 先加锁
		cond.L.Lock() //条件变量对应互斥加锁
		// 判断缓冲区是否满
		for len(out) == 5 {
			cond.Wait() // 1. 2. 3.		//挂起当前协程，等待条件变量满足，被消费者唤醒
		}
		num := rand.Intn(800) //产生一个随机数
		out <- num            //写到channel中
		fmt.Printf("生产者%dth，生产：%d\n", idx, num)
		cond.L.Unlock()                    // 访问公共区结束，并且打印结束，解锁
		cond.Signal()                      // 唤醒阻塞在条件变量上的 消费者
		time.Sleep(time.Millisecond * 200) //生产完休息一会，给其他协程执行机会
	}
}

func consumer08(in <-chan int, idx int) {
	for {
		// 先加锁
		cond.L.Lock() //条件变量对应互斥锁加锁
		// 判断 缓冲区是否为空
		for len(in) == 0 { //产品区为空，等待生产者生产
			cond.Wait() //挂起当前协程，等待条件变量满足，被生产者唤醒
		}
		num := <-in //将channel中的数据读走
		fmt.Printf("-----消费者%dth，消费：%d\n", idx, num)
		// 访问公共区结束后，解锁
		cond.L.Unlock() //消费结束，解除互斥锁
		// 唤醒 阻塞在条件变量上的 生产者
		cond.Signal()                      //唤醒阻塞的生产者
		time.Sleep(time.Millisecond * 200) //消费完休息一会，给其他协程执行机会
	}
}

func main() {
	product := make(chan int, 5)     //公共区使用channel模拟
	rand.Seed(time.Now().UnixNano()) //设置随机数种子
	quit := make(chan bool)          //创建用于结束通信的channel

	// 指定条件变量 使用的锁
	cond.L = new(sync.Mutex) // 创建互斥锁和条件变量

	for i := 0; i < 5; i++ {
		go producer08(product, i+1) // 1 生产者
	}
	for i := 0; i < 5; i++ {
		go consumer08(product, i+1) // 3 个消费者
	}
	/*	for {
		runtime.GC()
	}*/
	<-quit //主协程阻塞，不结束

}
