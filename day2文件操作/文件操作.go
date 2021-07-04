package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main01() {
	k := string("i love my work and i love myFamily too")
	a := strings.Split(k, "") //按指定分隔符进行拆分
	fmt.Println(a)
	a = strings.Fields(k)
	fmt.Println(a)                         //按空格进行拆分
	f := strings.HasSuffix("a.txt", ".tt") //判断字符串结束标记
	fmt.Println(f)
	f = strings.HasPrefix("a.txt", "a") //判断字符串起始标记
	fmt.Println(f)
}
func main02() {
	f, err := os.Create("E:/a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Println("successful")
}
func main03() {
	f, err := os.OpenFile("E:/a.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	_, err = f.WriteString("shjshj") //open没有写权限  写入失败
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("successful")
}

//按字符串写
func main04() {
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

/*fmt.Println(n)*/
func main05() {
	f, err := os.OpenFile("E:/a.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	defer f.Close()
	fmt.Println("打开文件成功")
	//创建一个带有缓冲区的reader
	reader := bufio.NewReader(f)
	for {
		buff, err := reader.ReadBytes('\n') //读一行数据
		fmt.Print(string(buff))
		if err == io.EOF {
			break
		}
	}
}
func main06() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("打印一个偶数", i)
			break
			fmt.Println("我不会被运行的")
		}
	}
}
func main07() {
	for i := 0; i < 10; i++ {
		if i > 5 {
			fmt.Println("执行")
			fmt.Println(i)
			break
		}
	}
	fmt.Println("只要break跳出循环我就会被运行")
}

//文件拷贝
func main() {
	//打开读文件
	fr, err := os.Open("G:/lfd.flv")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer fr.Close()
	//创建写文件
	fw, err := os.Create("E:/fz.flv")
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
