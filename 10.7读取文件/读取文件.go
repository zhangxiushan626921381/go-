package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main01() {
	//只读方式打开文件
	fp, err := os.Open("E:/a.txt")
	/*
		1.文件不存在
		2.文件权限
		3.文件打开上限
	*/
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	b := make([]byte, 9)
	//换行也会作为字符的一部分进行读取
	fp.Read(b)
	/*	fmt.Println(b)*/
	fmt.Println(string(b))
}
func main02() {
	fp, err := os.Open("E:/a.txt")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer fp.Close()
	//创建切片缓冲区
	r := bufio.NewReader(fp)
	//读取一行内容
	b, _ := r.ReadBytes('\n')
	//打印字符切片中ASCII值
	//将切片转成string打印就是汉字
	fmt.Println(string(b))
	//如果在文件截取中没有标志位或者分隔符就到文件末尾自动停止
	b, _ = r.ReadBytes('\n')
	fmt.Println(string(b))
}
func main03() {
	fp, err := os.Open("E:/a.txt")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer fp.Close()
	b := make([]byte, 20)
	for {
		_, err := fp.Read(b)
		//io.EOF表示文件的结尾  值为-1
		//eof end of file
		if err == io.EOF {
			break
		}
		fmt.Println(string(b))
	}
}
func main() {
	fp, err := os.Open("E:/a.txt")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer fp.Close()
	r := bufio.NewReader(fp)
	for {
		b, err := r.ReadBytes('\n')
		fmt.Println(string(b))
		if err == io.EOF {
			break
		}
		fmt.Println(string(b))
	}

}
