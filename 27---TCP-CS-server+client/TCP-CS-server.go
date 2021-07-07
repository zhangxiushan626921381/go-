package main

import (
	"fmt"
	"net"
)

func main() {
	//指定服务器通信协议，ip地址，端口号  创建一个用于监听的socket（套接字）
	listener, err := net.Listen("tcp", "127.0.0.1:8000") //需要用5000以上的
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}
	fmt.Println("服务器等待客户端建立连接......")
	defer listener.Close()
	//阻塞监听客户端连接请求  成功建立连接，返回用于通信的socket
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	fmt.Println("服务器与客户端成功连接")
	defer conn.Close()
	//读取客户端发送的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	conn.Write(buf[:n])
	//处理数据---打印
	fmt.Println("服务器端读到的数据：", string(buf[:n]))
	conn.Close()
}
