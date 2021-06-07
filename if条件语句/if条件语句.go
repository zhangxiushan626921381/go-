// if条件语句
package main

import (
	"fmt"
)

func main01() {
	var score int
	fmt.Println("请输入分数")
	fmt.Scan(&score)
	//格式 if 表达式 {代码体}

	if score > 700 {
		fmt.Println("我要上清华")
	} else if score > 650 {
		fmt.Println("我要上石油")
	} else {
		fmt.Println("我要回家种地")
	}
}
func main() {
	var a, b, c int
	fmt.Println("请输入三只小猪的体重")
	fmt.Scan(&a, &b, &c)
	if a > b {
		if a > c {
			fmt.Println("a重")
		} else {
			fmt.Println("c重")
		}
	} else {
		if a < b {
			if b > c {
				fmt.Println("b重")
			} else {
				fmt.Prin tln("c重")
			}
		}
	}
}
