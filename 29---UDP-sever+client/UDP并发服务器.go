package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//组织一个udp地址结构，指定服务器的ip+port
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr error:", err)
		return
	}
	fmt.Println("UDP服务器端地址结构建立成功")
	//创建用户通信的socket
	UdpCoon, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		fmt.Println("net.ListenUDP error", err)
		return
	}
	defer UdpCoon.Close()
	fmt.Println("UDP服务器通信socket建立成功")
	//读取客户端发送的数据
	buf := make([]byte, 4096)
	for {

		//返回3个值 分别是读取的字节个数，客户端的地址，error
		n, ClientAddr, err := UdpCoon.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("UdpCoon.ReadFromUDP error", err)
			return
		}
		//模拟处理数据
		fmt.Printf("服务器读到%v的数据:\n%s ", ClientAddr, string(buf[:n]))
		//提取系统当前时间
		daytime := time.Now().String()
		//回写数据给客户端
		_, err = UdpCoon.WriteToUDP([]byte(daytime), ClientAddr)
		if err != nil {
			fmt.Println("UdpCoon.WriteToUDP error:", err)
			return
		}

	}
}
