package main

import (
	"fmt"
	"net"
	"os"
)

func ReceiveFile(conn net.Conn, filename string) {
	//按照文件名创建新文件
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create error ", err)
		return
	}
	defer file.Close()
	//从网络中读数据写到本地文件
	buf := make([]byte, 4096)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("接收文件成功")
			return
		}
		//写入本地文件，读多少写多少
		file.Write(buf[:n])
	}

}
func main() {
	//创建用于监听的socket
	listener, err := net.Listen("tcp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("net.listen error ", err)
		return
	}
	defer listener.Close()
	//阻塞监听
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.accept error:", err)
		return
	}
	//获取文件名保存
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read error:", err)
		return
	}
	filename := string(buf[:n])
	//回写OK给发送端
	conn.Write([]byte("ok"))
	//获取文件内容
	ReceiveFile(conn, filename)
}
