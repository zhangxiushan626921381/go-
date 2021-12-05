package main

import "fmt"

type Transaction struct { //交易结构体
	//Id       int    //交易编号
	Sender   int //发送方
	Receiver int //接收方
}

/*type TransactionCollection struct { //交易集合结构体
	Tc  []int //交易集合切片
	num int      //该集合涉及的交易总数
}*/
var Tc []int
var Ts []int

var data []Transaction = make([]Transaction, 13)
var Tx []Transaction //交易结构体切片
//var TC []TransactionCollection //交易集合结构体切片
var Tcd map[string]int //交易集合与次数的键值对
var max int            //交易总数（具体基于交易频次的交易对数来定）
func Recursion(Tc []int) []int {
	//使用range
	/*for _, str := range data {
		if str != "" {
			out = append(out, str)
		}
	}*/
	//使用for循环遍历
	for i := 0; i <= max; i++ {
		if Tx[i].Sender == Tc[0] || Tx[i].Receiver == Tc[0] {
			Ts = append(Ts, Tx[i].Receiver)
			Ts = append(Ts, Tx[i].Sender)
		}
		if Tx[i].Sender == Tc[1] || Tx[i].Receiver == Tc[1] {
			Ts = append(Ts, Tx[max].Receiver)
			Ts = append(Ts, Tx[max].Sender)
		}
		Recursion(Tc)
	}
	return Ts
}
func main() {

	//造数据
	data = []Transaction{{1, 0}, {0, 2}, {0, 3}, {1, 4}, {1, 5},
		{2, 6}, {2, 7}, {3, 8}, {3, 9}, {4, 10},
		{4, 11}, {5, 12}, {5, 13}}
	//循环遍历交易数据   得到每一个分片的账户和交易总数
	Ts = make([]int, 0)
	Tc = []int{0, 1}
	Recursion(Tc)   //调用函数
	fmt.Println(Ts) //输出交易集合切片内容和次数

}
