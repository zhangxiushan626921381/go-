package main

import (
	"fmt"
)

type Transaction struct { //交易结构体
	Sender   int //发送方
	Receiver int //接收方
}

/*type TransactionCollection struct { //交易集合结构体
	Tc  []int   //交易集合切片
	//num int      //该集合涉及的交易总数
}*/
var data []Transaction = make([]Transaction, 100)

//var Tx []Transaction         //交易结构体切片
var Tc []int = make([]int, 100) //交易集合切片
var max int                     //交易数据的总数
var Ts []int = make([]int, 100)
var Td []int = make([]int, 100)
var a []int
var startlength, templength int

func Recursion(Tc []int) []int {
	max = 12
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

						Ts = append(Ts, data[i].Receiver)
						fmt.Println(Ts, data[i].Receiver)
						fmt.Println(len(Ts))

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
	data = []Transaction{{0, 2}, {0, 3},
		{1, 4}, {1, 5}, {2, 6}, {2, 7},
		{3, 8}, {3, 9}, {4, 10},
		{4, 11}, {5, 12}, {5, 13}}
	//var TC []TransactionCollection //交易集合结构体切片
	//Tx=make([]Transaction,15)
	Tc = []int{0, 1} //初始数据是0和1
	//循环遍历交易数据   得到每一个分片的账户和交易总数
	fmt.Println(data)
	Recursion(Tc)
	//var Ts []string //交易集合切片
	/*for i := 0; i < max; i++ {
		Recursion(Tc)                    //调用函数
		fmt.Println(TC[i].Tc) //输出交易集合切片内容和次数
	}*/
}
