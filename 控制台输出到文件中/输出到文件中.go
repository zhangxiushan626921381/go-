package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Create("E:/a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Println("successful")
	f, err = os.OpenFile("E:/a.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	defer f.Close()
	fmt.Println("打开文件成功")
	/*n,err:=f.WriteString("i love my family")   //默认从起始位置开始写的
	if err!=nil{
		fmt.Println(err)
		return*/
	off, _ := f.Seek(5, io.SeekCurrent)
	fmt.Println(off)
	n, _ := f.WriteAt([]byte("dshkkkhhh"), off)
	fmt.Println(n)
}
