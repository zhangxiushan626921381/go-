package main

import "fmt"

//定义结构体存储5名学生 三门成绩求出每名学生的总成绩和平均成绩
type student struct {
	id    int
	name  string
	score [3]int //结构体成员为数组  把3去掉就是切片

}

func main() {
	arr := []student{
		{101, "小明", [3]int{101, 102, 103}},
		{102, "小红", [3]int{104, 106, 108}},
		{103, "小刚", [3]int{100, 112, 123}},
		{104, "小强", [3]int{131, 152, 103}},
		{105, "小李", [3]int{106, 108, 108}}}
	fmt.Println(arr)
	//五名学生
	for i := 0; i < len(arr); i++ {
		//三门成绩
		sum := 0
		for j := 0; j < len(arr[i].score); j++ {
			sum += arr[i].score[j]
		}
		fmt.Printf("第%d个学生总成绩：%d ,平均分：%d\n", i+1, sum, sum/3)
	}

}
