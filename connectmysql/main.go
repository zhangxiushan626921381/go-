package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type Transaction struct {
	id       int
	Sender   int
	Receiver int
}

//连接数据库所需要的参数
const (
	USERNAME = "root"
	PASSWORD = "123456"
	IP       = "127.0.0.1"
	PORT     = "3306"
	dbName   = "LoginServer"
)

// DB 实例化一个数据库的连接池
var DB *sql.DB

// Init 写一个初始化数据库的方法
func Init() {
	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", IP, ":", PORT, ")/", dbName, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
	}
	fmt.Println("connect success!")

}
func InsertUser(TT Transaction) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail", err)
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO transactions (`id`,`Sender`,`Receiver`) VALUES (?,?,?)")
	if err != nil {
		fmt.Println("prepare fail", err)
		return false
	}
	//将参数传递到sql语句中并执行
	res, err := stmt.Exec(TT.id, TT.Sender, TT.Receiver)
	if err != nil {
		fmt.Println("Exec fail", err)
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}
func UpdateTransaction(TT Transaction) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail", err)
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE transactions SET Sender= ?,Receiver =? WHERE id =?")
	if err != nil {
		fmt.Println("prepare fail", err)
		return false
	}
	//将参数传递到sql语句中并执行
	res, err := stmt.Exec(TT.Sender, TT.Receiver, TT.id)
	if err != nil {
		fmt.Println("Exec fail", err)
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}
func DeleteTransaction(TT Transaction) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail", err)
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM transactions  WHERE id =?")
	if err != nil {
		fmt.Println("prepare fail", err)
		return false
	}
	//将参数传递到sql语句中并执行
	res, err := stmt.Exec(TT.id)
	if err != nil {
		fmt.Println("Delete Exec fail", err)
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

//查询
func SelectUserById(id int) {
	var user Transaction
	err := DB.QueryRow("SELECT * FROM transactions WHERE id = ?", id).Scan(&user.id, &user.Sender, &user.Receiver)
	if err != nil {
		fmt.Println("查询出错了", err)
	}
	fmt.Println(user)
}
func main() {
	Init()
	Tx1 := Transaction{id: 1, Sender: 16, Receiver: 91}
	/*Tx2 := Transaction{2, 20, 27}
	Tx3 := Transaction{id: 3, Sender: 19, Receiver: 81}*/
	/*	Tx4:=Transaction{id:4,Sender:5,Receiver: 6}
		InsertUser(Tx4)
		/*InsertUser(Tx1)
		InsertUser(Tx2)
		InsertUser(Tx3)*/
	/*UpdateTransaction(Tx1)
	UpdateTransaction(Tx2)
	UpdateTransaction(Tx3)
	SelectUserById(1)
	SelectUserById(2)
	SelectUserById(3)*/
	DeleteTransaction(Tx1)
	SelectUserById(1)
}
