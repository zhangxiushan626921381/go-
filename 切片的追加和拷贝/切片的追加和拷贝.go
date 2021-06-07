package main

import (
	"fmt"
)

func main01() {
	var s []int
	//s[0]=123
	//在使用append添加数据时 切片的地址可能会发生变化 如果容量扩充导致数据存储不下，
	//切片会自动找寻新的空间存储数据 同时释放之前的数据
	//使用append添加数据
	//添加数据默认是偶数，如果是2,caps就是2，如果是3,caps就是4，如果是5就是6
	s = append(s, 12, 23, 56, 56, 56)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Printf("%p\n", s)
	s = append(s, 12, 23, 56, 56, 56)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Printf("%p\n", s)
}
func main() {
	s := []int{1, 2, 5, 6, 7}
	s1 := make([]int, 3)
	//将s中的内容复制到s1中，s1中农要有足够的容量，
	//使用拷贝操作后s s1是两个独立的空间不会相互影响
	copy(s1, s)
	s1[2] = 123
	fmt.Println("s1=", s1)
}

//

func main03() {
	s1 := []int{1, 2, 15, 1, 5}
	s2 := make([]int, 5)
	//如果想拷贝具体的切片中的片段，需要使用切片中的截取
	copy(s2, s1[1:4])
	fmt.Println(s2)

}
