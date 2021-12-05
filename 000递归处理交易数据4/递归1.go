package main

import (
	"fmt"
)

type Transaction struct { //交易结构体
	Sender   int //发送方
	Receiver int //接收方
}

var data []Transaction = make([]Transaction, 100)
var Tc []int = make([]int, 100) //交易集合切片
var Ts []int = make([]int, 100)
var Td []int = make([]int, 100)
var startlength, templength, max int

func Recursion(Tc []int) []int {
	max = 16 //改变数据记得改总数
	startlength = 0
	Ts = Tc
	templength = len(Ts)
	fmt.Println(Ts)
	for h := 0; h <= 1; h++ {
		if startlength < templength {
			for j := startlength; j < templength; j++ {
				//找关联的账号 第一句
				for i := 0; i < max; i++ {
					//有的话Tx[i].Receiver和Tx[i].Sender不加
					if data[i].Sender == Tc[j] && data[i].Receiver != Tc[j] {
						if data[i].Receiver != 0 && data[i].Receiver != 1 {
							Ts = append(Ts, data[i].Receiver)
							fmt.Println(Ts, data[i].Receiver)
							fmt.Println(len(Ts))
						}
					} else {
						if data[i].Receiver == Tc[j] && data[i].Sender != Tc[j] {
							if data[i].Sender != 0 && data[i].Sender != 1 {
								Ts = append(Ts, data[i].Sender)
								fmt.Println(Ts, data[i].Sender)
								fmt.Println(len(Ts))
							}
						} else {
							if data[i].Receiver == Tc[j] && data[i].Sender == Tc[j] {
								fmt.Println(Ts)
							}
						}
					}
				}
			}
		}
		Tc = Ts
		startlength = len(Tc) - (len(Tc) - 2)
		templength = len(Ts)
	}
	return Ts
}
func main() {
	//造数据
	data = []Transaction{{0, 2}, {3, 9}, {0, 3},
		{1, 4}, {1, 5}, {2, 6},
		{2, 7}, {3, 8}, {36, 2},
		{4, 10}, {16, 0}, {13, 19},
		{4, 11}, {5, 12}, {5, 13},
		{5, 66}, {0, 14}, {56, 3}}
	//var TC []TransactionCollection //交易集合结构体切片
	//Tx=make([]Transaction,15)
	Tc = []int{0, 1} //初始数据是0和1
	//循环遍历交易数据   得到每一个分片的账户和交易总数
	fmt.Println(data, len(data))
	Recursion(Tc)
}
