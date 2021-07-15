package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	//获取用户键盘输入（stdin），将输入数据发送给服务器
	go func() {
		str := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				continue
			}
			//写给服务器
			conn.Write(str[:n])
		}
	}()
	//接收服务器返回的大写数据
	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("监测到服务器关闭，客户端也关闭")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("服务器返回的数据是", string(buf[:n]))
	}

}
