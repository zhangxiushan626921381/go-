package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func httpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1 //封装函数内部的错误，传出给调用者
		return
	}
	defer resp.Body.Close()
	//循环读取网页的数据，传出给调用者
	for {
		buf := make([]byte, 4096)
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n]) //累加每一次循环读到的buf数据，存入到result一次性返回
	}
	return
}

//爬取页面操作
func working(start int, end int) {
	fmt.Printf("正在爬取第%d页到第%d页的内容\n", start, end)
	//循环爬取页面的内容
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=steam&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		result, err := httpGet(url)
		if err != nil {
			fmt.Println("heepGet error", err)
			continue //继续打印其他的页面
		}
		//fmt.Println("result=:",result)
		//将读到的网页数据保存成一个文件
		file, err := os.Create("第" + strconv.Itoa(i) + "页" + ".html")
		if err != nil {
			fmt.Println("Create err:", err)
			continue
		}
		file.WriteString(result)
		file.Close() //保存好一个文件关闭一个文件
	}
}
func main() {
	//指定爬虫起始终止页
	var start, end int
	fmt.Println("请输入爬取的起始页：")
	fmt.Scan(&start)
	fmt.Println("请输入爬取的终止页：")
	fmt.Scan(&end)
	//封装函数
	working(start, end)
}
