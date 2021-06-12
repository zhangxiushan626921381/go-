package main

import "fmt"

type student struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}

func main() {
	s := student{101, "貂蝉", "男", 19, "徐州"}
	//结构体变量赋值
	s1 := s
	s1.sex = "女"
	fmt.Println(s)
	fmt.Println(s1)
	//比较两个结构体是否相同  使用==可以对结构体成员进行比较操作
	if s1 == s {
		fmt.Println("相同")
	} else {
		fmt.Println("不相同")
	}
	//大于小于可以比较结构体成员  比如年龄
	if s1.age > s.age {
		fmt.Println("s1 big")
	} else {
		fmt.Println("s big")
	}
}
