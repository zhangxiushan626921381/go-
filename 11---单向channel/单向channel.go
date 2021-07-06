package main

import "fmt"

func main01() {
	ch := make(chan int) //双向channel
	var sendCh chan<- int = ch
	sendCh <- 789
	/*	num:=<-sendCh*/ //
	var recvCh <-chan int = ch
	num := <-recvCh
	fmt.Println(num)
}
func send(out chan<- int) {
	out <- 789
	close(out)
}
func recv(in <-chan int) {
	n := <-in
	fmt.Println("读到一个值：", n)
}
func main() {
	var ch chan int
	ch = make(chan int) //双向channel
	go func() {
		send(ch) //双向channel转为单向channel
	}()

	recv(ch)

}
