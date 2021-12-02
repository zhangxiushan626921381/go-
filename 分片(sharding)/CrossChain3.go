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
	ShardNum := 6 //分片的数量
	TxNum := 100  //交易的总数
	UserNum := 12 //用户的数量
	var UserNumInChain int = UserNum / ShardNum
	var Tx []Transaction
	Tx = make([]Transaction, 101)
	var CrossChainCount = 0
	var Shard string
	//生产交易
	var a1, a2, a3, a4, a5, a6 int
	/*	var b1, b2, b3, b4, b5, b6 int*/
	var b1 int
	//rand.Seed(time.Now().UnixNano()) //随机数种

	for i := 0; i < TxNum; i++ { //循环生成交易
		//TransAction := make([][]int, 100001)

		a1 = rand.Intn(10)
		a2 = rand.Intn(10) + 10
		a3 = rand.Intn(30) + 20
		a4 = rand.Intn(30) + 50
		a5 = rand.Intn(10) + 80
		a6 = rand.Intn(10) + 90

		if b1 == a1 {
			Tx[TxNum].Sender = rand.Intn(2) //随机生成发送方
			Tx[TxNum].Receiver = rand.Intn(12)
			if Tx[TxNum].Sender == Tx[TxNum].Receiver { //如果相等则重新生成接收方
				Tx[TxNum].Receiver = rand.Intn(12)
			}
		} else {
			if b1 == a2 {
				Tx[TxNum].Sender = rand.Intn(2) + 2 //随机生成发送方
				Tx[TxNum].Receiver = rand.Intn(12)
				if Tx[TxNum].Sender == Tx[TxNum].Receiver { //如果相等则重新生成接收方
					Tx[TxNum].Receiver = rand.Intn(12)
				}
			} else {
				if b1 == a3 {
					Tx[TxNum].Sender = rand.Intn(2) + 4 //随机生成发送方
					Tx[TxNum].Receiver = rand.Intn(12)
					if Tx[TxNum].Sender == Tx[TxNum].Receiver { //如果相等则重新生成接收方
						Tx[TxNum].Receiver = rand.Intn(12)
					}
				} else {
					if b1 == a4 {
						Tx[TxNum].Sender = rand.Intn(2) + 6 //随机生成发送方
						Tx[TxNum].Receiver = rand.Intn(12)
						if Tx[TxNum].Sender == Tx[TxNum].Receiver { //如果相等则重新生成接收方
							Tx[TxNum].Receiver = rand.Intn(12)
						}
					} else {
						if b1 == a5 {
							Tx[TxNum].Sender = rand.Intn(2) + 8 //随机生成发送方
							Tx[TxNum].Receiver = rand.Intn(12)
							if Tx[TxNum].Sender == Tx[TxNum].Receiver { //如果相等则重新生成接收方
								Tx[TxNum].Receiver = rand.Intn(12)
							}
						} else {
							if b1 == a6 {
								Tx[TxNum].Sender = rand.Intn(2) + 10 //随机生成发送方
								Tx[TxNum].Receiver = rand.Intn(12)
								if Tx[TxNum].Sender == Tx[TxNum].Receiver { //如果相等则重新生成接收方
									Tx[TxNum].Receiver = rand.Intn(12)
								}
							}
						}
					}
				}
			}
		}

		//a3 = rand.Intn(30) + 20

		//a4 = rand.Intn(30) + 50

		//a5 = rand.Intn(10) + 80

		//a6 = rand.Intn(10) + 90

		//随机生成接收方

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
		fmt.Printf("随机数%d,发送方%d，接收方%d，发送分片%d，接收分片%d,判断:%s\n", b1, Tx[TxNum].Sender, Tx[TxNum].Receiver, Tx[TxNum].SenderShard, Tx[TxNum].ReceiverShard, Shard)
		//fmt.Println(CrossChainCount) //跨链次数
		//fmt.Printf("发送方%d接收方%d\n", Tx[TxNum].Sender, Tx[TxNum].Receiver)
		/*var jiaoyixinxi []int
		jiaoyixinxi = append(jiaoyixinxi, Tx[TxNum].Sender)
		fmt.Println(jiaoyixinxi)*/
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
