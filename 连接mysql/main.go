package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type User struct {
	Username string
	Password string
}

//连接数据库所需要的参数
const (
	USERNAME = "root"
	PASSWORD = "123456"
	IP       = "127.0.0.1"
	PORT     = "3306"
	dbName   = "loginserver"
)

//实例化一个数据库的连接池
var DB *sql.DB

func main() {
	InitDB()
}

//写一个初始化数据库的方法
func InitDB() {
	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", IP, ":", PORT, ")/", dbName, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail err", err)
		return
	}
	fmt.Println("connect success!")
}
