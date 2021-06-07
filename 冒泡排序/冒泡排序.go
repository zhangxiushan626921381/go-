package main

import "fmt"

func main() {
	arr := [10]int{10, 9, 16, 7, 8, 6, 12, 5, 15, 1}
	//外层控制行
	for i := 0; i < len(arr)-1; i++ {
		//内层控制列
		for j := 0; j < len(arr)-1-i; j++ {
			//比较两个相邻元素满足条件 交换数据
			//升序排序使用大于号 降序排序使用小于号
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)

}
