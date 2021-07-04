package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main01() {
	//获取用户输入的目录路径
	fmt.Println("请输入待查询的目录")
	var path string
	fmt.Scan(&path)
	//打开目录
	fp, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	//读取目录项
	info, err := fp.Readdir(-1) // -1 代表读取目录中的所有目录项
	//遍历返回的切片
	for _, fileinfo := range info {
		if fileinfo.IsDir() {
			fmt.Println(fileinfo.Name(), "是一个目录")
		} else {
			fmt.Println(fileinfo.Name(), "是一个文件")
		}

	}
}
func main02() {
	//获取用户输入的目录路径
	fmt.Println("请输入待查询的目录")
	var path string
	fmt.Scan(&path)
	//打开目录
	fp, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	//读取目录项
	info, err := fp.Readdir(-1) // -1 代表读取目录中的所有目录项
	//遍历返回的切片
	for _, fileinfo := range info {
		if !fileinfo.IsDir() {
			if strings.HasSuffix(fileinfo.Name(), ".jpg") {
				fmt.Println("jpg文件有：", fileinfo.Name())
			}
		}
	}
}

//拷贝指定文件到指定目录
func copyDir(path1 string, path2 string) {
	//打开读文件
	fr, err := os.Open(path1)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer fr.Close()
	//创建写文件
	fw, err := os.Create(path2)
	if err != nil {
		fmt.Println("创建文件失败", err)
		return
	}
	defer fw.Close()
	//从读文件中获取数据，放到缓冲区中
	buf := make([]byte, 1024)
	//循环从读文件中获取数据，写到写文件中
	for {
		n, err := fr.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("读完", n)
			return
		}
		fw.Write(buf[:n]) //读多少 写多少
	}

}
func main() {
	//获取用户输入的目录路径
	fmt.Println("请输入待查询的目录")
	var path string
	fmt.Scan(&path)
	//打开目录
	fp, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	//读取目录项
	info, err := fp.Readdir(-1) // -1 代表读取目录中的所有目录项
	//遍历返回的切片
	for _, fileInfo := range info {
		if !fileInfo.IsDir() {
			if strings.HasSuffix(fileInfo.Name(), ".jpg") {
				fmt.Println("jpg文件有：", fileInfo.Name())
				copyDir(path+"/"+fileInfo.Name(), "./"+fileInfo.Name())
			}
		}
	}
}

//统计指定目录内单词出现的个数
//根据用户指定的目录打开
//找到目录中的.txt文件（有可能有多个）
//打开其中一个文件，循环读取一行//reader:=bufio.NewReader,reader.ReadBytes('\n')
//将一行数据的字符串拆分存到切片中，split,fileds
//遍历切片，统计单词出现的次数
