package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
		}
		close(ch) //写端写完数据主动关闭channel
	}()
	/*for {
		if num,ok:=<-ch;ok==true{
			fmt.Println("读到数据",num)
		}else{
			k:=<-ch
			fmt.Println("关闭后读取数据",k)   //读到默认值0
		break
		}
	}*/
	for num := range ch {
		fmt.Println("读到数据：", num)
	}
}
