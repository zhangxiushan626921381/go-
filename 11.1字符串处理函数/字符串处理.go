package main

import (
	"fmt"
	"strings"
)

//Contains
func main01() {
	str := "hello world"
	//在一个字符串中查找另外一个字符串是否出现  模糊查找（搜索手机出来一系列关于手机的词语）返回值为布尔类型
	value := strings.Contains(str, "llo") //在str中查找llo是否出现
	if value == true {
		fmt.Println("出现")
	} else {
		fmt.Println("没出现")
	}
	fmt.Println(value)
}

//join 功能：字符串连接
func main02() {
	s := []string{"1235", "5656", "4556", "5454"}
	//将一个字符串切片拼接成一个字符串
	str := strings.Join(s, "---")
	fmt.Println(str)
}

//index 功能：在字符串s中找sep所在的位置  如果找不到返回-1
func main03() {
	str := "日照香炉生紫烟"
	//在一个字符串中查找另外一个字符是否出现，返回值为整型
	i := strings.Index(str, "找香炉")
	fmt.Println(i)
}

//Repeat 功能：重复s字符串count次，返回重复完成的字符串
func main04() {
	str := "蚂蚁和大象大象和蚂蚁"
	//将一个字符串重复n次 n取值>=0
	value := strings.Repeat(str, 3)
	fmt.Println(value)
}

//Replace 功能：在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
func main05() {
	str := "hello world,我操你大爷的"
	//屏蔽关键字
	b1 := strings.Replace(str, "操", "*", 3)
	fmt.Println(b1)
}

//split 功能：把s字符串按照sep分割，返回slice
func main06() {
	str := "hello.world.ha"
	value := strings.Split(str, ".")
	fmt.Println(value)
}

//Trim 功能：在s字符串的头部和尾部去除cutset指定的字符串
func main07() {
	str := "    hello.world.ha   "
	//去掉字符串前后的内容
	value := strings.Trim(str, " ")
	fmt.Println(value)
}

//Fields 功能：去除s字符串的空格符，并且按照空格分割返回slice
func main() {
	str := "  hi hello world  sb "
	value := strings.Fields(str)
	fmt.Println(value)
}
