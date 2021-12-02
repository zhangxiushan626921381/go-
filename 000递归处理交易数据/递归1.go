package main

import "fmt"

type Transaction struct { //交易结构体
	Id       int    //交易编号
	Sender   string //发送方
	Receiver string //接收方
}
type TransactionCollection struct { //交易集合结构体
	Tc  []string //交易集合切片
	num int      //该集合涉及的交易总数
}

var Tx []Transaction           //交易结构体切片
var TC []TransactionCollection //交易集合结构体切片
var Tcd map[string]int         //交易集合与次数的键值对
var max int                    //交易总数（具体基于交易频次的交易对数来定）
func Recursion(Tc []string) []string {
	//使用range
	/*for _, str := range data {
		if str != "" {
			out = append(out, str)
		}
	}*/
	//使用for循环遍历
	for i := 0; i <= max; i++ {
		if Tx[i].Sender == Tc[0] || Tx[i].Receiver == Tc[0] {
			Tc = append(Tc, Tx[i].Receiver)
			Tc = append(Tc, Tx[i].Sender)
		}
		if Tx[i].Sender == Tc[1] || Tx[i].Receiver == Tc[1] {
			Tc = append(Tc, Tx[max].Receiver)
			Tc = append(Tc, Tx[max].Sender)
		}
		Recursion(Tc)
	}
	return Tc
}
func main() {
	//循环遍历交易

	var Ts []string //交易集合切片
	for i := 0; i < max; i++ {
		Recursion(Ts)
		fmt.Println(TC[i].Tc)
	}
}
