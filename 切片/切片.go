package main

import "fmt"

func main01() {
	//切片定义  var 切片名 []数据类型
	//var s [] int
	//自动推导类型创建切片  make([]数据类型，5)
	a := make([]int, 5)
	//通过下标为切片赋值
	a[0] = 123
	a[1] = 566
	a[2] = 659
	a[3] = 856
	a[4] = 666
	//a[6]=665//err
	//通过append添加切片信息
	a = append(a, 665, 4656, 56)
	fmt.Println(a)
	//通过len查看切片长度
	fmt.Println(len(a))
}
func main02() {
	a := make([]int, 5)
	a[0] = 123
	a[1] = 566
	a[2] = 659
	a[3] = 856
	a[4] = 666
	//遍历
	/*for i:=0;i<len(a);i++{
	fmt.Println(a[i])*/
	for i, v := range a {
		fmt.Println(i, v)
	}
}
func main03() {
	//不写元素个数叫做切片，必须写元素个数叫做数组
	var s []int = []int{1, 2, 1, 6, 5}
	s = append(s, 5, 5, 56)
	//容量大于等于长度
	fmt.Println("长度：", len(s))
	fmt.Println("容量：", cap(s))
	//容量每次扩展为上次的倍数
	s = append(s, 5, 5, 56)
	fmt.Println("长度：", len(s))
	fmt.Println("容量：", cap(s))
}
func main() {
	a := make([]int, 5)
	a[0] = 1
	a = append(a, 454)
	//再切片打印时，只能打印有效长度中的数据 cap(a)不能作为数据打印的条件
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])

	}
}
