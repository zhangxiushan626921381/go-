package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//文件拷贝 从一个文件中拷贝数据到新的文件中
	fp1, err1 := os.Open("E:\\硕鼠\\区块.flv")
	fp2, err2 := os.Create("G:/lfd.flv")
	if err1 != nil || err2 != nil {
		fmt.Println("拷贝失败")
		return
	}
	defer fp1.Close()
	defer fp2.Close()
	b := make([]byte, 1024)
	for {
		//块读取
		n, err := fp1.Read(b)
		if err == io.EOF {
			break
		}
		fp2.Write(b[:n])
	}
	fmt.Println("拷贝成功")
}
