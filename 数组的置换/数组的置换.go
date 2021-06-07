package main

import (
	"fmt"
)

func main() {
	var arr [10]int = [10]int{1, 15, 5, 12, 6, 8, 7, 16, 9, 10}
	//定义起始和结束下标
	start := 0
	end := len(arr) - 1
	for {
		if start > end {
			break
		}
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	fmt.Println(arr)
}
