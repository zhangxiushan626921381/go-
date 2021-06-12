package main

import "fmt"

type student struct {
	id    int
	name  string
	age   int
	sex   string
	score int
	addr  string
}

func main01() {
	//var 结构体变量名 结构体类型
	//定义结构体数组
	//var 结构体数组名 [元素个数]结构体类型
	var arr [5]student
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i].id, &arr[i].name, &arr[i].age, &arr[i].sex, &arr[i].score, &arr[i].addr)
	}

	//结构体排序  可以根据结构体成员进行排序
	/*
		101 宿 15 女 15 宿
		102 谁 26 方 15 是
		103 是 23 是 23 是
		107 搜 56 啥 65 是
		108 算 56 是 56 是

	*/
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			//比较结构体成员年龄
			if arr[j].age > arr[j+1].age {
				//结构体数组中的元素   允许相互赋值
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
func main() {
	//[元素个数]数组    []切片len
	arr := []student{{101, "曹操", 48, "男", 90, "许昌"},
		{101, "夏侯惇", 40, "男", 85, "荆州"},
		{101, "张辽", 45, "男", 89, "许昌"}}
	//在切片中添加数据
	arr = append(arr, student{104, "许褚", 39, "男", 99, "许昌"})
	for _, v := range arr {
		fmt.Println(v)
	}
}
