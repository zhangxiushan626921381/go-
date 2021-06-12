package main

import "fmt"

func main() {
	m := make(map[int]string, 1)
	//在Map中值允许重复
	m[0] = "张飞"
	m[1] = "刘备"
	m[2] = "张飞"
	//在Map打印中如果没有出现定义的map  默认定义为空
	//	fmt.Println("------",m[3],"------")
	for v, k := range m {
		fmt.Println(v, k)
	}
	v, ok := m[3]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("没有m[3]这个值")
	}
	//fmt.Println(m)
}
