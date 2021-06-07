package main

import "fmt"

func main0() {
	//map
	//var map1 map[int]string //定义map1 map 这是一个空map
	map1 := make(map[int]string, 1)
	//map的容量和长度是自动扩容的 容量和长度相等
	//map[key] key是一个基本的数据类型
	//map中的数据时无序存储的
	map1[1] = "张三" //这里面的1并不是下标，可以是任何数字
	map1[100] = "王五"
	map1[5] = "贺荣"
	//遍历
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	fmt.Println(map1)
}
func main1() {
	m := make(map[string]int)
	m["贺荣"] = 3
	m["贺荣2"] = 4
	fmt.Println(m)
}
func main() {
	//map在定义时key是唯一的，不允许重复

	//map中key和value不可以反过来操作

}
