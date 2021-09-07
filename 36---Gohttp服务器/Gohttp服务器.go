package main

import (
	"fmt"
	"net/http"
)

func myHandle(w http.ResponseWriter, r *http.Request) {
	//w:写给客户端的数据内容
	//r:从客户端读到的内容
	w.Write([]byte("this is a web server"))
	fmt.Println("header:", r.Header)
	fmt.Println("url:", r.URL)
	fmt.Println("host:", r.Host)
	fmt.Println("remoteAddr:", r.RemoteAddr)
	fmt.Println("method:", r.Method)
}
func main() {
	//注册回调函数，该函数在客户端访问服务器时会被自动调用
	http.HandleFunc("/itcast", myHandle)
	//绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}
