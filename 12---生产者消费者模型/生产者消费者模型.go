package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 8; i++ {
		fmt.Println("生产数据", i)
		out <- i
	}
	close(out)
}
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("消费者拿到数据", num)
		time.Sleep(2 * time.Second)
	}

}
func main() {
	ch := make(chan int)
	go producer(ch) //子go程扮演生产者
	consumer(ch)    //主go程扮演消费者
}
