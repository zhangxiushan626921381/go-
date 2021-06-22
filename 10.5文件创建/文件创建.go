package main

import (
	"fmt"
	"os"
)

func main() {
	//路径分为绝对路径（从盘符开始到文件的位置 linux中是以/目录开始的）和相对路径（相对某一个文件的位置）
	//\\是一个转义字符 在一般文件操作中都使用/
	fp, err := os.Create("E:/a.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return //如果return出现在主函数中表示程序的结束
	}
	fmt.Println("文件创建成功")
	//关闭文件
	//1.占用内存和缓冲区
	//2.文件打开上限
	defer fp.Close()
}
