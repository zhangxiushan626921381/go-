package main

import "fmt"

type stu struct {
	id    int
	name  string
	score int
	age   int
}

func test(st stu) {
	st.name = "李逵"
	fmt.Println(st)
	fmt.Println(&st.name)

}
func main01() {
	st := stu{103, "宋江", 98, 56}
	//结构体作为函数参数 值传递
	test(st)
	fmt.Println(st)
	fmt.Println(&st.name)

}

//所有的切片都属于地址传递
func test1(stu []stu) {
	stu[0].name = "晁盖"
}
func main02() {
	//结构体切片
	stus := []stu{{101, "宋江", 9, 89},
		{102, "卢俊义", 9, 89}}
	//为切片添加信息
	stus = append(stus, stu{103, "吴用", 8, 78})
	//结构体切片作为函数参数是地址传递
	test1(stus)
	fmt.Println(stus)
}
func test2(stu [2]stu) {
	stu[0].name = "晁盖"
}
func main() {
	//结构体数组
	stus := [2]stu{{101, "宋江", 9, 89},
		{102, "卢俊义", 9, 89}}

	test2(stus)
	fmt.Println(stus)
}
