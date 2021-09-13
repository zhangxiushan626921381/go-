package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func sum(k string) map[string]int {
	value := strings.Fields(k) //将读到的数据拆分成字符串切片
	m := make(map[string]int)  //创建一个用于存储交易信息出现次数的map
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
	fmt.Println("请输入路径")
	var path string
	fmt.Scan(&path)
	f, err := os.OpenFile(path, os.O_RDONLY, 6)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	fmt.Println("文件打开成功")
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		buff, _ := reader.ReadBytes('.') //读数据
		fmt.Print(string(buff))
		m := sum(string(buff))
		//h为交易信息  g为交易次数
		for h, g := range m {
			fmt.Println(h, g)
		}
	}
}
