package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("正在唱歌---sing")
		time.Sleep(100 * time.Millisecond)
	}
}
func dance() {
	for i := 0; i < 5; i++ {
		fmt.Println("正在跳舞---dance")
		time.Sleep(100 * time.Millisecond)
	}

}
func main() {
	go sing()
	go dance()
	for {

	}
}
