package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strings"
)

type Transaction struct { //定义交易的结构体
	Id            int //数据编号
	Sender        int //发送方
	Receiver      int //接收方
	SenderShard   int //发送分片
	ReceiverShard int //接收分片
	IsCrossChain  int //是否跨链
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
	stmt, err := tx.Prepare("INSERT INTO transactions (`id`,`Sender`,`Receiver`,`SenderShard`,`ReceiverShard`,`IsCrossChain`) VALUES (?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("prepare fail", err)
		return false
	}
	//将参数传递到sql语句中并执行
	res, err := stmt.Exec(TT.Id, TT.Sender, TT.Receiver, TT.SenderShard, TT.ReceiverShard, TT.IsCrossChain)
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

func main() {
	Init()
	ShardNum := 10 //分片的数量
	TxNum := 1000  //交易的总数
	UserNum := 100 //用户的数量
	var UserNumInChain int = UserNum / ShardNum
	var Tx []Transaction
	Tx = make([]Transaction, 100001)
	var CrossChainCount = 0
	var Shard string
	//生产交易
	for i := 0; i < TxNum; i++ { //循环生成交易
		//TransAction := make([][]int, 100001)
		//rand.Seed(time.Now().Uni xNano()) //随机数种子
		Tx[TxNum].Sender = rand.Intn(99)            //随机生成发送方
		Tx[TxNum].Receiver = rand.Intn(99)          //随机生成接收方
		if Tx[TxNum].Sender == Tx[TxNum].Receiver { //如果相等则重新生成接收方
			Tx[TxNum].Receiver = rand.Intn(99)
		}
		Tx[TxNum].SenderShard = Tx[TxNum].Sender / UserNumInChain
		Tx[TxNum].ReceiverShard = Tx[TxNum].Receiver / UserNumInChain

		if Tx[TxNum].ReceiverShard != Tx[TxNum].SenderShard {
			Tx[TxNum].IsCrossChain = 1
			Shard = "0"
			CrossChainCount = CrossChainCount + 1
		} else {
			Tx[TxNum].IsCrossChain = 0
			Shard = "1"
		}
		//输出全部交易信息
		fmt.Printf("发送方%d，接收方%d，发送分片%d，接收分片%d,判断:%s\n", Tx[TxNum].Sender, Tx[TxNum].Receiver, Tx[TxNum].SenderShard, Tx[TxNum].ReceiverShard, Shard)
		//fmt.Println(CrossChainCount) //跨链次数
		//fmt.Printf("发送方%d接收方%d\n", Tx[TxNum].Sender, Tx[TxNum].Receiver)

		InsertUser(Tx[TxNum])
	}
	/*var SameShard int
	if Tx[TxNum].SenderShard ==Tx[TxNum].ReceiverShard {
		SameShard++
	}*/
	/*fmt.Println(ShardNum)*/
	//计算跨链率
	var CrossChainRate float64
	CrossChainRate = float64(CrossChainCount) / float64(TxNum)
	fmt.Println(CrossChainRate)
	fmt.Println(TxNum)
}
