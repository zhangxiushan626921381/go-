package main

import (
	"fmt"
	"strconv"
)

func main01() {
	//将其他类型转换成字符串
	//将bool类型转成字符串
	s := strconv.FormatBool(false)
	fmt.Println(s)
	//FormatInt(数据，进制)
	//将整型转成字符串
	n := strconv.FormatInt(15, 2)
	fmt.Println(n)
	//将浮点型转成字符串
	//FormatFloat(数据,'f',保留小数位数（会四舍五入）,位数(64/32)
	k := strconv.FormatFloat(1.235, 'f', 2, 64)
	fmt.Println(k)
}

//将字符串转成整型
func main02() {
	str := "false1"
	//只能将“false”“true”转成布尔类型  多余的数据是无效的
	a, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}
}

//将字符串转成整型
func main03() {
	str := "56568989"
	a, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

//将字符串转成浮点型
func main04() {
	str := "1.2365"
	a, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
func main() {
	b := make([]byte, 0, 1024)
	b = strconv.AppendBool(b, false)
	b = strconv.AppendInt(b, 123, 10)
	b = strconv.AppendFloat(b, 1.236, 'f', 5, 64)
	b = strconv.AppendQuote(b, "hello")
	fmt.Println(string(b))
}
