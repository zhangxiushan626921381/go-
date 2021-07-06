package main

import (
	"fmt"
	"sync"
	"time"
)

//使用channel完成同步
/*var ch=make(chan int)
func printer(str string){
	for _,c:=range str{
		fmt.Printf("%c",c)
		time.Sleep(time.Millisecond*3)
	}
}
func person1(){				//先执行
	printer("hello")
	ch<-1
}
func person2(){				//后执行
	<-ch
	printer("你好")


}
func main01()  {
	go person2()
	go person1()
	for{
		;
	}
}*/
//使用锁完成同步
var mutex sync.Mutex //创建一个互斥锁  新建的互斥锁状态为0 未加锁  锁只有一把
func printer(str string) {
	mutex.Lock() //访问共享数据前加锁
	for _, c := range str {
		fmt.Printf("%c", c)
		time.Sleep(time.Millisecond * 3)
	}
	mutex.Unlock() //共享数据访问结束解锁
}
func person1() { //先执行
	printer("hello")

}
func person2() { //后执行

	printer("你好")

}
func main() {
	go person1()
	go person2()
	for {

	}
}
