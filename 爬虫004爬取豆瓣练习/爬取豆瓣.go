package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

//爬取指定的页面返回result
func httpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	//循环爬取整页数据
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}
func SaveFile(i int, all [][]string) {
	file, err := os.Create("第" + strconv.Itoa(i) + "页.txt")
	if err != nil {
		fmt.Println("os.Create error", err)
		return
	}
	defer file.Close()
	n := len(all)
	fmt.Println(n)
	for k := 0; k < n; k++ {
		file.WriteString(all[k][1])
	}
}
func spider(i int) {
	url := "https://tieba.baidu.com/f?kw=steam&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	result, err := httpGet(url)
	if err != nil {
		fmt.Println("httpGet err", err)
		return
	}
	//fmt.Println("result:",result)
	//使用正则表达式过滤需要的信息并保存到文件
	//解析编译正则表达式
	reg := regexp.MustCompile(`title="主题作者:(?s:(.*?))"`)
	//提取需要的信息
	all := reg.FindAllStringSubmatch(result, -1)
	//将提取的信息封装到函数中
	SaveFile(i, all)

}
func ToWork(start int, end int) {
	fmt.Printf("正在爬取第%d页到%d页的内容\n", start, end)
	for i := start; i <= end; i++ {
		spider(i)
	}

}
func main() {
	var start, end int
	fmt.Println("请输入爬取的起始页(>=1)")
	fmt.Scan(&start)
	fmt.Println("请输入爬取的终止页(>=start)")
	fmt.Scan(&end)
	//封装爬取函数
	ToWork(start, end)
}
