package main

import "fmt"

func main() {
	max := 0
	var arr [10]int = [10]int{2, 5, 6, 8, 8, 4, 65, 656, 8, 57}
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}

	}
	fmt.Println("最重的小猪体重：", max)
}
