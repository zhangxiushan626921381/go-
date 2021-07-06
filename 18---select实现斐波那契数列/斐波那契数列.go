package main

import (
	"fmt"
	"runtime"
)

func Fbnq(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Println(num)
		case <-quit:
			runtime.Goexit()
		}
	}
}
func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go Fbnq(ch, quit) //子go程打印FBNQ数列
	x := 1
	y := 1
	for i := 0; i < 20; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- false
}
