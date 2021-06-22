package main

import "fmt"

func main() {
	arr := make([]interface{}, 5)
	arr[0] = 123
	arr[1] = 1.23
	arr[2] = "hello world"
	arr[3] = [3]int{1, 2, 3}
	for _, v := range arr {
		//对数据v进行类型断言
		if data, ok := v.(int); ok {
			fmt.Println("整型数据", data)
		} else if data, ok := v.(float64); ok {
			fmt.Println("浮点型数据", data)
		} else if data, ok := v.(string); ok {
			fmt.Println("字符串数据", data)
		} else if data, ok := v.([3]int); ok { //必须写3
			fmt.Println("数组数据", data)
		}
	}
}
