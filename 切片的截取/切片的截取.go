package main

import "fmt"

func main() {
	s := []int{1, 2, 5, 5, 4}
	//切片名[起始下标:]
	slice := s[2:]
	//切片名[:结束位置] 不包含结束位置
	slice1 := s[:2]
	//切片名[起始位置：结束位置]
	slice2 := s[0:2]
	//切片名[起始位置：结束位置：容量]容量要写小于切片的值
	slice3 := s[2:4:5]
	fmt.Println(slice)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(cap(slice3))

}
