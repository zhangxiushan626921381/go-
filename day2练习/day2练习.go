package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*1.根据用户指定的目录打开
2.找到目录中的.txt文件（有可能有多个），
3.打开其中一个文件，循环读取一行reader:=bufio.NewReader,reader.ReadBytes('\n')
4.将一行数据的字符串拆分存到切片中，split,fileds
5.遍历切片，统计单词出现的次数*/
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
	//获取用户输入的目标路径
	fmt.Println("请输入路径")
	var path string
	fmt.Scan(&path)
	OF, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("文件打开失败，错误是", err)
		return
	}
	defer OF.Close()
	//读取目录项
	info, err := OF.Readdir(-1) // -1 代表读取目录中的所有目录项
	//遍历返回的切片
	for _, file := range info {
		if !file.IsDir() {
			if strings.HasSuffix(file.Name(), ".txt") {
				fmt.Println(".txt文件有:", file.Name())
			}
		}
		f, err := os.OpenFile(path+"/"+file.Name(), os.O_RDONLY, 6)
		if err != nil {
			fmt.Println("文件打开失败", err)
			return
		}
		fmt.Println("文件打开成功")
		defer f.Close()
		reader := bufio.NewReader(f)
		for {
			buff, _ := reader.ReadBytes('\n') //读一行数据
			fmt.Print(string(buff))
			m := sum(string(buff))
			for h, g := range m {
				fmt.Println(h, g)
			}
		}
	}
}
