package main

import "fmt"

type student struct {
	id    int
	name  string
	age   int
	score int
}

func main01() {
	//定义map
	m := make(map[int]student)
	m[101] = student{102, "周瑜", 32, 101}
	m[102] = student{102, "孙策", 35, 112}
	//打印数据
	fmt.Println(m[102])
	//打印数据格式
	fmt.Printf("%T\n", m[101])
	//遍历数据
	for i, v := range m {
		fmt.Println(i, v)
	}
}
func main() {
	m := make(map[int][]student)
	m[101] = append(m[101],
		student{102, "孙策", 35, 112},
		student{103, "孙权", 37, 192},
		student{104, "周瑜", 39, 162},
		student{104, "太史慈", 39, 162})
	m[102] = append(m[102],
		student{103, "刘备", 40, 192},
		student{103, "张飞", 44, 192},
		student{103, "关羽", 85, 192})
	m[103] = append(m[103],
		student{103, "曹丕", 56, 192},
		student{103, "曹植", 63, 192},
		student{103, "曹彰", 22, 192})
	for i, v := range m {
		//遍历map中的信息
		fmt.Println(i, v)
	}
	for i, v := range m { //遍历map中的信息
		for k, data := range v {
			fmt.Println("key:", i, k, data)
		}
	}
}
