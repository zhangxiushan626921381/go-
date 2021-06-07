package main

import (
	"fmt"
)

func test(arr [10]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main01() {
	var arr [10]int = [10]int{1, 5, 1, 8, 8, 8: 10}
	//数组作为函数参数是传递数组名
	//数组作为函数参数是值传递的
	test(arr)

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
	var arr [10]int = [10]int{1, 5, 1, 8, 8, 8: 10}
	BubbleSort(arr)
	arr = BubbleSort(arr)
	fmt.Println(arr)

}
