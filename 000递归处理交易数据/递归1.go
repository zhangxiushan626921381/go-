package main

type Transaction struct {
	Id       int    //交易编号
	Sender   string //发送方
	Receiver string //接收方
}

var Tx []Transaction
var Tc []string
var Tb, b []string

var max int //交易总数（具体基于交易频次的交易对数来定）
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
	return b
}
func main() {

}
