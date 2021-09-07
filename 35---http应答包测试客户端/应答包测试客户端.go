package main

import (
	"fmt"
	"net"
)

func errFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
	}
}

//装浏览器
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	errFunc(err, "net.Dial error:")
	defer conn.Close()
	httpRequest := "GET /itcast HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n"
	conn.Write([]byte(httpRequest))
	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	if n == 0 {
		return
	}
	fmt.Println(string(buf[:n]))
}
