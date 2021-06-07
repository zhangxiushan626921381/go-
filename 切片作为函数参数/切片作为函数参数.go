package main

import "fmt"

func test(s []int) {
	//形参和实参使用的是同一内存地址空间 地址传递
	//地址传递  形参能够改变实参的值
	s[2] = 123
	fmt.Println(s)
}
func main01() {
	s := []int{1, 2, 3, 4, 5}
	test(s)
}

//思想   冒泡排序  每次一定可以确定一个数的位置
func BubbleSort(s []int) {
	//决定冒泡排序的次数
	for i := 0; i < len(s); i++ {
		//从i位置开始冒泡排序确定一个数的最终位置
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]

			}
		}
	}
}
func main02() {
	s := []int{9, 1, 5, 6, 7}
	//在实际开发中  建议实用切片来代替数组
	s = append(s, 3, 10, 2, 4, 8)
	BubbleSort(s)
	fmt.Println(s)
}
func test011(s []int) {
	//一般情况下，不要在函数中使用append操作
	//如果使用append操作切片可能改变切片的地址 需要使用返回值给实参赋值
	s = append(s, 4, 5, 6)
	fmt.Printf("%p\n", s)
	fmt.Println(s)

}
func main03() {
	s := []int{1, 2, 3}
	fmt.Printf("%p\n", s)
	test1(s)
	fmt.Println(s)
}
func test1(s []int) []int { //返回值是[]int
	//一般情况下，不要在函数中使用append操作
	s = append(s, 4, 5, 6)
	fmt.Printf("%p\n", s)
	fmt.Println(s)
	return s
}
func main() {
	s := []int{1, 2, 3}
	fmt.Printf("%p\n", s)
	s = test1(s) //新的空间更细 覆盖s
	fmt.Println(s)
	fmt.Printf("%p\n", s)
}
