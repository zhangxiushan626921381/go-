package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//获取随机数种子
	rand.Seed(time.Now().UnixNano())
	//生成一个100-999之间的随机数
	//random:=rand.Intn(900)+100
	//生成一个100-999之间的随机数
	random1 := make([]int, 3)
	random1[0] = rand.Intn(9) + 1 //1-9
	random1[1] = rand.Intn(10)
	random1[2] = rand.Intn(10)
	usernum := make([]int, 3)
	var num int
	for {
		fmt.Println("请输入一个三位数")
		fmt.Scan(&num)
		if num >= 100 && num <= 999 {
			break
		}
		fmt.Println("输入错误，请重新输入")
	}
	usernum[0] = num / 100
	usernum[1] = num / 10 % 10
	usernum[2] = num % 10
	for i := 0; i < 3; i++ {

		if usernum[i] > random1[i] {
			fmt.Printf("您输入的第%d位数太大了\n", i+1)
		} else if usernum[i] < random1[i] {
			fmt.Printf("您输入的第%d位数太小了\n", i+1)
		} else {
			fmt.Printf("恭喜你，第%d位数相同\n", i+1)

		}

	}

	fmt.Println(random1)
	fmt.Println(num)
}
