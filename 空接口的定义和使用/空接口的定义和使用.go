package main

import "fmt"

func main01() {
	//空接口  可以接收任意类型的数据
	var i interface{}
	i = 10
	fmt.Println(i)
	fmt.Printf("%p\n", &i)
	i = "hello world"
	fmt.Println(i)
	fmt.Printf("%p\n", &i)
}
func main() {
	//空接口切片
	var i []interface{}
	i = append(i, 1, 2, "hello world", "你看啥", [3]int{1, 23})
	for j := 0; j < len(i); j++ {
		fmt.Println(i[j])
	}
	for i, v := range i {
		fmt.Println(i, v)
	}
}
