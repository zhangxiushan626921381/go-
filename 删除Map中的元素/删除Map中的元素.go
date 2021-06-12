package main

import "fmt"

func main() {
	m := map[int]string{101: "刘备", 102: "张飞", 103: "关羽"}
	//打印m对应的类型
	//fmt.Printf("%T",m)
	//删除 delete(map，key)
	delete(m, 103)
	//delete在删除map中的元素的时候如果key不存在也不会报错
	delete(m, 105)
	fmt.Println(m)
}
