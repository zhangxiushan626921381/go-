package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("生产者生产一个数字", i)
		out <- i
	}
	close(out)
}
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("消费者获得一个数字", num)
		time.Sleep(time.Millisecond * 300)
	}
}
func main() {
	product := make(chan int)
	go producer(product)
	go consumer(product)

	for {

	}
}
