package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendfile(conn net.Conn, filepath string) {
	//只读打开文件
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer file.Close()
	//从本地文件中读取数据写给接收端，读多少，写多少
	buf := make([]byte, 4096)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("file.Read error:", err)
			}
			return
		}
		//写到网络socket中
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn.Write error", err)
		}
	}
}
func main() {
	list := os.Args //获取命令行参数

	if len(list) != 2 {
		fmt.Println("格式为:go run xxx.go 文件绝对路径")
		return
	}
	//提取绝对路径
	filepath := list[1]
	//提取文件名
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("os.Stat error", err)
		return
	}
	filename := fileInfo.Name()
	//filesize:=fileInfo.Size()
	//主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("net.Dial error", err)
		return
	}
	defer conn.Close()
	//发送文件名给接收端
	_, err = conn.Write([]byte(filename))
	if err != nil {
		fmt.Println("conn.Write error", err)
		return
	}
	//读取服务器回发的ok
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read error", err)
		return
	}
	//写文件内容给服务器---借助conn
	if "ok" == string(buf[:n]) {
		sendfile(conn, filepath)
	}

}
