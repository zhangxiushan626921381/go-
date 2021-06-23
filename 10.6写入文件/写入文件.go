package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//os.Create()创建文件时，文件不存在会创建一个新文件，如果文件存在会覆盖源内容
	fp, err := os.Create("E:/a.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	fmt.Println("文件创建成功")
	defer fp.Close()
	//写入文件
	n, _ := fp.WriteString("rainbowIT\r\n")
	fmt.Println(n)
	//在go语言中一个汉字是3个字符
	n, _ = fp.WriteString("大象")
	fmt.Println(n)
	fp.WriteString("蚂蚁")
}
func main02() {
	fp, err := os.Create("E:/a.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	//关闭文件
	defer fp.Close()
	//将字符切片写入文件中
	/*b:=[]byte{'d','d','f','d'}
	//文件对象.write(字符切片)
	fp.Write(b)*/
	//将字符串转成字符切片写入文件中
	//str:="hello world"
	str := "锄禾日当午"
	//字符串和字符切片允许转换
	b := []byte(str)
	fp.Write(b)
}
func main03() {
	//fp,err:=os.Create("E:/a.txt")
	//打开文件
	//os.open(文件名)只读方式打开文件
	//os.Open("E:/a.txt")
	//os.openfile(文件名，打开模式，打开权限)
	//openfile不能创建新文件

	fp, err := os.OpenFile("E:/a.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer fp.Close()
	//获取文件的字符个数
	//负数向左偏移  正数向右偏移
	n, _ := fp.Seek(2, io.SeekStart) //后面是偏移位置
	fmt.Println(n)
	//通过字符串获取字符切片
	b := []byte("蚂蚁和大象")
	//当使用writeAt进行指定位置插入数据时  会依次覆盖原内容
	fp.WriteAt(b, n) // 后面是偏移位置
}
