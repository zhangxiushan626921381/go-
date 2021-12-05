package main

/*
import "fmt"

type Transaction struct { //交易结构体
	//Id       int    //交易编号
	Sender   int //发送方
	Receiver int //接收方
}

type TransactionCollection struct { //交易集合结构体
	Tc  []string //交易集合切片
	num int      //该集合涉及的交易总数
}

var Tx []Transaction         //交易结构体切片
//造数据  ...
var TC []TransactionCollection //交易集合结构体切片
var max int                    //交易总数（具体基于交易频次的交易对数来定）
var Tc []string
func Recursion(Tc []string) []string {
	//使用range
	/*for _, str := range data {
	   if str != "" {
	      out = append(out, str)
	   }
	}*/
//使用for循环遍历
/*j:=0
	Tc.legth:=1
	//for j := Tc.legth-1; j <= Tc.legth; j++{

	Tc[0]=0;Tc[1]=1;
	startlegth=0;templegth=2;
	while（startlegth<templegth）{

		//记住下次的开始
		startlegth=templegth;
		for(j := startlegth; j <= templegth; i++)
		{
		//找关联的账号 第一句
		for i := 0; i <= max; i++ {
		if Tx[i].Sender == Tc[j] || Tx[i].Receiver == Tc[j] {
		//有的话Tx[i].Receiver和Tx[i].Sender不加
		Tc = append(Tc, Tx[i].Receiver)
		Tc = append(Tc, Tx[i].Sender)
		}
		}
		j++;
		}//根据上一次找出的关联账号集合找下一次关联账号集合



		templegth=Tc.legth;

	}

}
func main() {
	//循环遍历交易数据   得到每一个分片的账户和交易总数
	var Ts []string //交易集合切片
	for i := 0; i < max; i++ {
		Recursion(Ts)                    //调用函数
		fmt.Println(TC[i].Tc, TC[i].num) //输出交易集合切片内容和次数
	}
}
*/
