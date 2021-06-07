package main

import "fmt"

func main01() {
	//一维数组
	//var 数组名 [元素个数] 数据类型
	//二维数组
	//var 数组名 [行个数] [列个数] 数据类型
	var arr [3][4]int = [3][4]int{{1, 2, 8}, {8, 9, 6, 9}, {5, 6}}
	//通过行和列找到具体元素

	//fmt.Println(arr)
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}
}
func main() {
	var arr [3][4]int
	//len(二维数组名) 二维数组的行数
	fmt.Println(len(arr))
	//arr[0]表示二位数组的第一行
	//len(arr[下标]) 二维数组的列数
	fmt.Println(len(arr[0]))
}
