package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("i:", i)
		time.Sleep(1 * time.Second) // 延时1秒
	}
}
func main01() {
	//创建一个goroutine，启动另外一个任务
	go newTask()
	i := 0
	//main goroutine 循环打印
	for {
		i++
		fmt.Println("i:", i)
		time.Sleep(1 * time.Second) // 延时1秒
	}

}
func main() {
	go newTask()
	fmt.Println("main 运行结束退出")
}
func main03() {
	go func() { //创建一个子Go程
		for i := 0; i < 5; i++ {
			fmt.Println("i'm goroutine")
			time.Sleep(time.Second)
		}
	}()
	for i := 0; i < 5; i++ { //主Go程
		fmt.Println("i'm main")
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}

}
