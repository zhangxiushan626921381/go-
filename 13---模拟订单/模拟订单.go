package main

import "fmt"

type OrderInfo struct { //创建订单结构体
	id int
}

func producer(out chan<- OrderInfo) { //定义生成订单
	for i := 0; i < 5; i++ { //循环生成5份订单
		num := OrderInfo{i + 1}
		fmt.Println("提交一个订单", num)
		out <- num //写入仓库
	}
	close(out) //写完关闭仓库 不关闭的话消费者会一直等待
}
func consumer(in <-chan OrderInfo) { //处理订单
	for k := range in { //从仓库中取出订单
		fmt.Println("收到订单号为：", k) //模拟处理订单
	}
}
func main() {
	ch := make(chan OrderInfo) //定义一个订单仓库
	go producer(ch)            //建立协程，传只写channel
	consumer(ch)               //建立协程，传只读channel
}
