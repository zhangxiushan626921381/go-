package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// Client 创建用户结构体
type Client struct {
	C    chan string
	Name string
	Addr string
}

// OnlineMap 创建全局Map,存储在线用户
var OnlineMap map[string]Client

//创建全局channel 传递用户消息
var message = make(chan string)

func WriteMsgToClient(conn net.Conn, clnt Client) {
	//监听用户自带channel上是否都有消息
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}
func MakeMsg(clnt Client, msg string) (buf string) {
	buf = "[" + clnt.Addr + "]" + clnt.Name + ":" + msg
	return
}
func handleConnect(conn net.Conn) {
	defer conn.Close() //go程结束 socket也结束
	//获取用户的网络地址 IP+port
	netAddr := conn.RemoteAddr().String() //RemoteAddr返回远端网络地址
	//创建新连接用户的结构体信息,默认用户名是IP+port
	Clnt := Client{make(chan string), netAddr, netAddr}
	//将新连接的用户添加到map中   key:ip+port value:结构体信息
	OnlineMap[netAddr] = Clnt
	//创建专门给当前用户发送消息的go程
	go WriteMsgToClient(conn, Clnt)
	//发送用户上线到message通道中
	message <- MakeMsg(Clnt, "login")
	//创建一个channel，用来判断用户退出状态
	isquit := make(chan bool)
	//创建一个channel 判断用户是否活跃
	hasdata := make(chan bool)
	//创建匿名go程 专门处理用户发送的消息
	go func() {
		for {
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if n == 0 {
				isquit <- true
				fmt.Printf("检测到用户%s退出\n", Clnt.Name)
				return
			}
			if err != nil {
				fmt.Println("conn.Read error:", err)
				return
			}
			//将读到的用户消息写入到message通道中
			msg := string(buf[:n-1])

			//提取在线用户列表
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("Online user list:\n"))
				//遍历当前map，获取在线用户
				for _, users := range OnlineMap {
					userInfo := users.Addr + ":" + users.Name + "\n"
					conn.Write([]byte(userInfo))
				}
				//判断用户发送的改名命令
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				newname := strings.Split(msg, "|")[1]
				Clnt.Name = newname //修改结构体成员name
				newaddr := strings.Split(msg, "|")[2]
				Clnt.Addr = newaddr       //修改结构体成员addr
				OnlineMap[netAddr] = Clnt //更新onlinemap列表
				conn.Write([]byte("rename successful\n"))
			} else {
				message <- MakeMsg(Clnt, msg)
			}
			hasdata <- true
		}
	}()
	//保证不退出
	for {
		//监听channe上的数据流动
		select {
		case <-isquit: //用户不主动退出 阻塞
			//先删除当前用户然后再去广播 然后再去删除相关的go程
			delete(OnlineMap, Clnt.Addr)       //将用户从online移除
			message <- MakeMsg(Clnt, "logout") //写入用户退出消息到全局channel 进行广播
			return                             //结束当前用户对应的协程
		case <-hasdata: //select有一个分支可走  就会回去循环不会触发下面的case
		//什么都不做  目的是重置下面case的计时器
		case <-time.After(time.Second * 4):
			delete(OnlineMap, Clnt.Addr)       //将用户从online移除
			message <- MakeMsg(Clnt, "logout") //写入用户退出消息到全局channel 进行广播
			return

		}
	}
}
func Manager() {
	//初始化OnlineMap
	OnlineMap = make(map[string]Client)
	//循环从message中获取消息
	for {
		//循环监听全局channel中是否有数据,有数据存储到msg，无数据就阻塞
		msg := <-message
		//循环发送消息给所有的在线用户 要想执行必须 msg:=message 执行完，解除阻塞
		for _, clnt := range OnlineMap {
			clnt.C <- msg
		}
	}
}
func main() {
	//创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.listen error:", err)
		return
	}
	defer listener.Close()
	//创建管理者go程
	go Manager()
	//循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error:", err)
			return
		}
		//启动go程处理客户端的请求
		go handleConnect(conn)
	}
}
