package main

import "fmt"

func main() {
	//指定爬取起始终止页数
	var start, end int
	fmt.Println("请输入爬取的起始页（>=1):")
	fmt.Scan(&start)
	fmt.Println("请输入爬取的终止页（>=start):")
	fmt.Scan(&end)

}
