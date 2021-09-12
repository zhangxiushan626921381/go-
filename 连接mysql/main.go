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
	id       int
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
	/*user1:=User{"zhang","123123",1}
	InsertUser(user1)*/
	//DeleteUser(user1)
	/*user2:=User{"he","45678",1}
	UpdateUser(user2)*/
	SelectUserById(1)
}

//查询
func SelectUserById(id int) {
	var user User
	err := DB.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Username, &user.Password, &user.id)
	if err != nil {
		fmt.Println("查询出错了", err)
	}
	fmt.Println(user)
}
func UpdateUser(user User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE users SET name = ?, password = ? WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.Username, user.Password, user.id)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}
func DeleteUser(user User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.id)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	//获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}
func InsertUser(user User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO users (`name`, `password`,`id`) VALUES (?, ?,?)")
	if err != nil {
		fmt.Println("Prepare fail", err)
		return false
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(user.Username, user.Password, user.id)
	if err != nil {
		fmt.Println("InsertUser Exec fail err", err)
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
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
