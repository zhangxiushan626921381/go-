package main

import (
	"fmt"
	"strings"
)

func main01() {
	var m1 map[int]string //声明map，没有空间，不能直接存储key---value
	//m1[1]="s"
	fmt.Println(m1)
	m2 := map[int]string{1: "s", 2: "f"}
	fmt.Println(m2)
	m3 := make(map[int]string)
	m3[1] = "s"
	fmt.Println(m3)
	m4 := make(map[int]string, 10)
	m4 = map[int]string{1: "s", 2: "f"}
	fmt.Println(m4)
	var m5 map[int]string = map[int]string{1: "lufei"}
	fmt.Println(m5)
	var s1 []int = []int{1, 2, 3}
	fmt.Println(s1)
}
func main02() {
	s2 := map[int]string{1: "s", 2: "f"}
	for k, v := range s2 {
		fmt.Println(k, v)
	}
}
func main03() {
	s := map[int]string{1: "s", 2: "f"}
	if v, h := s[1]; !h {
		fmt.Println(v, h)
	} else {
		fmt.Println(v, h)
	}
}
func mapdelete(m map[int]string, key int) map[int]string {
	delete(m, 2)
	return m
}
func main04() {
	m := map[int]string{1: "ss", 2: "dd", 366: "fg"}
	m2 := mapdelete(m, 2)
	fmt.Println(m2)
}
func sum(k string) map[string]int {
	value := strings.Fields(k) //将字符串拆分成字符串切片
	m := make(map[string]int)  //创建一个用于存储字符出现次数的map
	for i := 0; i < len(value); i++ {
		if _, ok := m[value[i]]; ok == true { // ok==true 说明value[i]这个key存在
			m[value[i]]++
		} else { //ok==false 说明这个key不存在  添加到value
			m[value[i]] = 1
		}
	}
	return m
}
func main() {
	k := string("i love my work and i love myFamily too")
	l := sum(k)
	//遍历map 展示每个word出现的次数
	for h, g := range l {
		fmt.Println(h, g)
	}
}
