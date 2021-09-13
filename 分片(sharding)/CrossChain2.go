package main

import (
	"fmt"
	"math/rand"
)

func main() {
	type Transaction struct { //定义交易的结构体
		Sender        int //发送方
		Receiver      int //接收方
		SenderShard   int //发送分片
		ReceiverShard int //接收分片
		IsCrossChain  int //是否跨链
	}
	ShardNum := 9  //分片的数量
	TxNum := 10000 //交易的总数
	UserNum := 100 //用户的数量
	var UserNumInChain int = UserNum / ShardNum
	var Tx []Transaction
	Tx = make([]Transaction, 10001)
	var CrossChainCount = 0
	//var Shard string
	//生产交易
	for i := 0; i < TxNum; i++ { //循环生成交易
		//TransAction := make([][]int, 100001)
		//rand.Seed(time.Now().UnixNano()) //随机数种子
		Tx[TxNum].Sender = rand.Intn(99)            //随机生成发送方
		Tx[TxNum].Receiver = rand.Intn(99)          //随机生成接收方
		if Tx[TxNum].Sender == Tx[TxNum].Receiver { //如果相等则重新生成接收方
			Tx[TxNum].Receiver = rand.Intn(99)
		}
		Tx[TxNum].SenderShard = Tx[TxNum].Sender / UserNumInChain
		Tx[TxNum].ReceiverShard = Tx[TxNum].Receiver / UserNumInChain

		if Tx[TxNum].ReceiverShard != Tx[TxNum].SenderShard {
			Tx[TxNum].IsCrossChain = 1
			//Shard="跨分片"
			CrossChainCount = CrossChainCount + 1
		} else {
			Tx[TxNum].IsCrossChain = 0
			//Shard="不跨分片"
		}
		//输出全部交易信息
		//fmt.Printf("发送方%d，接收方%d，发送分片%d，接收分片%d,判断:%s\n", Tx[TxNum].Sender, Tx[TxNum].Receiver, Tx[TxNum].SenderShard, Tx[TxNum].ReceiverShard,Shard)
		/*fmt.Println(CrossChainCount)*/ //跨链次数
		fmt.Printf("发送方%d接收方%d\n", Tx[TxNum].Sender, Tx[TxNum].Receiver)
		var jiaoyixinxi []int
		jiaoyixinxi = append(jiaoyixinxi, Tx[TxNum].Sender)
		fmt.Println(jiaoyixinxi)
	}

	/*var SameShard int
	if Tx[TxNum].SenderShard ==Tx[TxNum].ReceiverShard {
		SameShard++
	}*/
	/*fmt.Println(ShardNum)*/
	//计算跨链率
	//var CrossChainRate float64
	//CrossChainRate=float64(CrossChainCount)/float64(TxNum)
	//fmt.Println(CrossChainRate)
	fmt.Println(TxNum)
}
