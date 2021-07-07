package main

import (
	"fmt"
	"net"
	"strings"
)

func handlerConnect(conn net.Conn) {
	defer conn.Close()
	//获取连接的客户端addr
	addr := conn.RemoteAddr()
	fmt.Println("客户端成功连接！", addr)
	//循环读取客户端发送数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		fmt.Println(buf[:n])
		if "exit\r\n" == string(buf[:n]) || "exit\n" == string(buf[:n]) {
			fmt.Println("服务器接收到exit请求，关闭服务器")
			return
		}
		if n == 0 {
			fmt.Println("服务器监测到客户端已经关闭，断开连接")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("服务器读到：", string(buf[:n]))
		//完成小写转大写，回发给客户端
		strings.ToUpper(string(buf[n]))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
func main() {
	//创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}
	defer listener.Close()
	//创建客户端连接请求
	for {
		fmt.Println("服务器等待客户端连接.....")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}
		//具体完成服务器和客户端的通信数据
		go handlerConnect(conn)
	}
}
