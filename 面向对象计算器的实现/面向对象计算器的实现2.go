package main

import "fmt"

type fulei struct {
	num1 int
	num2 int
}
type jiafazilei struct {
	fulei
}
type jianfazilei struct {
	fulei
}
type Jiekou interface {
	jisuan() int
}

func (add *jiafazilei) jisuan() int {
	return add.num1 + add.num2
}
func (less *jianfazilei) jisuan() int {
	return less.num1 - less.num2
}
func Duotai(jiekou Jiekou) (value int) {
	value = jiekou.jisuan()
	return value
}

type Gongchang struct {
}

func (gongchang *Gongchang) jisuanjieguo(num1 int, num2 int, yunsuanfu string) (value int) {
	//创建接口变量
	var jiekou Jiekou
	switch yunsuanfu {
	case "+":
		var adjf jiafazilei = jiafazilei{fulei{num1, num2}}
		//value=adjf.jisuan()
		jiekou = &adjf
	case "-":
		var lsjf jianfazilei = jianfazilei{fulei{num1, num2}}
		//value=lsjf.jisuan()
		jiekou = &lsjf
	}
	//value=jiekou.jisuan()
	value = Duotai(jiekou)
	return value
}

func main() {
	var gongchang Gongchang
	value := gongchang.jisuanjieguo(10, 20, "-")
	fmt.Println(value)
}
