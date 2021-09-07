package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	type transaction struct { //定义交易的结构体
		Sender        int //发送方
		Receiver      int //接收方
		SenderShard   int //发送分片
		ReceiverShard int //接收分片
		IsCrossChain  int //是否跨链
	}

	ShardNum := 10  //分片的数量
	TxNum := 100000 //交易的总数
	UserNum := 100  //用户的数量
	var UserNumInChain int = UserNum / ShardNum
	var Tx transaction
	Tx.Sender = rand.Intn(99)
	Tx.Receiver = rand.Intn(99)
	Tx.SenderShard = Tx.Sender / UserNumInChain
	Tx.ReceiverShard = Tx.Receiver / UserNumInChain
	//生产交易
	for i := 0; i < TxNum; i++ { //循环生成交易
		rand.Seed(time.Now().UnixNano()) //随机数种子
		Tx.Sender = rand.Intn(99)        //随机生成发送方
		Tx.Receiver = rand.Intn(99)      //随机生成接收方
		if Tx.Sender == Tx.Receiver {    //如果相等则重新生成接收方
			Tx.Receiver = rand.Intn(99)
		}
	}
	var CrossChainCount int = 0
	for i := 0; i <= TxNum; i++ {
		if Tx.ReceiverShard != Tx.SenderShard {
			Tx.IsCrossChain = 1
			CrossChainCount++
		} else {
			CrossChainCount = 0
		}
		fmt.Println(CrossChainCount)
	}
}
