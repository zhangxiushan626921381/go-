package main

import (
	"fmt"
	"net/http"
)

func main() {
	//获取服务器应答包内容
	res, err := http.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println("http.GET error:", err)
		return
	}
	defer res.Body.Close()
	//简单查看应答包
	fmt.Println("header", res.Header)
	fmt.Println("status", res.Status)
	fmt.Println("statusCode", res.StatusCode)
	fmt.Println("proto", res.Proto)
	buf := make([]byte, 4096)
	var result string
	for {
		n, err := res.Body.Read(buf)
		if n == 0 {
			fmt.Println("read finish")
			return
		}
		if err != nil {
			fmt.Println("Body.read err:", err)
			break
		}
		result += string(buf[:n])
	}
	fmt.Printf("|%v|\n", result)
}
