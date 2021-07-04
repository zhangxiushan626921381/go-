package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3) //存满3个元素之前不会阻塞
	fmt.Println("len(ch)", len(ch), "cap(ch)", cap(ch))
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
			fmt.Println("子Go程i", i, "len(ch)", len(ch))

		}

	}()
	time.Sleep(1 * time.Second)
	for i := 0; i < 5; i++ {
		num := <-ch
		time.Sleep(1 * time.Second)
		fmt.Println("主Go程读到", num)

	}
}
