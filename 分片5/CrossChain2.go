package main

import (
	"fmt"
	"math/rand"
)
type Transaction1 struct { //定义交易的结构体
	Id            int //数据编号
	Sender        int //发送方
	Receiver      int //接收方
	SenderShard   int //发送分片
	ReceiverShard int //接收分片
	IsCrossChain  int //是否跨链
}

func main() {

	ShardNum := 6 //分片的数量
	TxNum := 100 //交易的总数
	UserNum := 12  //用户的数量
	var UserNumInChain int = UserNum / ShardNum
	var Tx []Transaction1
	Tx = make([]Transaction1, 101)
	var CrossChainCount = 0
	var Shard string
	//生产交易
	var b1 int
	//rand.Seed(time.Now().UnixNano()) //随机数种
	for i := 0; i < TxNum; i++ { //循环生成交易
		//TransAction := make([][]int, 100001)
		b1 = rand.Intn(100)
		if b1>=0&&b1<10 {
			Tx[TxNum].Sender = 4  //随机生成发送方
			Tx[TxNum].Receiver = 5
			Tx[TxNum].IsCrossChain = 1
			Shard = "跨分片"
			CrossChainCount = CrossChainCount + 1
		} else {
			if b1>=10&&b1<20 {
				Tx[TxNum].Sender = 10 //随机生成发送方
				Tx[TxNum].Receiver = 11
				Tx[TxNum].IsCrossChain = 1
				Shard = "跨分片"
				CrossChainCount = CrossChainCount + 1
			} else {
				if b1>=20&&b1<100  {
					Tx[TxNum].Sender = rand.Intn(12)  //随机生成发送方
					Tx[TxNum].Receiver = rand.Intn(12)
					if Tx[TxNum].Sender == Tx[TxNum].Receiver { //如果相等则重新生成接收方
						Tx[TxNum].Receiver = rand.Intn(12)
					}
				}
			}
		}
		Tx[TxNum].SenderShard = Tx[TxNum].Sender / UserNumInChain
		Tx[TxNum].ReceiverShard = Tx[TxNum].Receiver / UserNumInChain
		if Tx[TxNum].ReceiverShard != Tx[TxNum].SenderShard {
			Tx[TxNum].IsCrossChain = 1
			Shard = "跨分片"
			CrossChainCount = CrossChainCount + 1
		} else {
			Tx[TxNum].IsCrossChain = 0
			Shard = "不跨分片"
		}
		//输出全部交易信息
		fmt.Printf("随机数%d,发送方%d，接收方%d，发送分片%d，接收分片%d,判断:%s\n", b1,Tx[TxNum].Sender, Tx[TxNum].Receiver, Tx[TxNum].SenderShard, Tx[TxNum].ReceiverShard, Shard)
		fmt.Println(CrossChainCount) //跨链次数

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