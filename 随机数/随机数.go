package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main01() {
	//导入头文件 math/rand time
	//添加随机数种子
	//使用随机数
	start := time.Now().Nanosecond()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		//根据时间 1970.1.1.0.0.0随机的
		fmt.Println(rand.Intn(100))
		end := time.Now().Nanosecond()
		fmt.Println(end - start)
	}
}
func BubbleSort(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr

}
func main() {
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(rd.Intn(100))
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(50)
	}
	arr = BubbleSort(arr)
	fmt.Println(arr)
}

//福利彩票  红球6个 1-33 不能重复  蓝球是1-16 随机一个
//rand.Intn(33)+1//1-33
