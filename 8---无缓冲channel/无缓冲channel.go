package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子go程，i:", i)
			ch <- i
		}
	}()

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		num := <-ch
		fmt.Println("主Go程读到的数据", num)
	}

}
