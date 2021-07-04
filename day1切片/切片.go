package main

import "fmt"

func main01() {
	var a []int = []int{1, 2, 3, 4, 6, 9}
	b := a[1:3:5]
	c := a[0:3]
	fmt.Println(b)
	fmt.Println(len(b)) //3-1==2
	fmt.Println(cap(b)) //5-1==4
	fmt.Println(cap(c)) //6-0==6
}
func main02() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a := arr[2:5] //{3,4,5} 左闭右开
	fmt.Println(a)
	fmt.Println(len(a)) //5-2==3
	fmt.Println(cap(a)) //10-2==8
	a1 := a[2:7]
	fmt.Println(a1)      //{5,6,7,8,9}
	fmt.Println(len(a1)) //7-2==5
	fmt.Println(cap(a1)) //8-2==6
}
func main03() {
	arr := make([]int, 5)
	arr = []int{1, 2, 3}
	fmt.Println(arr)
}
func noempty(data []string) []string {
	out := data[:0] //在原切片上截取一个长度为0的切片
	for _, str := range data {
		if str != "" {
			out = append(out, str)
		} //取到空字符串无作为
	}
	return out
}
func main04() {
	data := []string{"red", "", "black", "", "", "pink", "blue"}
	afterdata := noempty(data)
	fmt.Println(afterdata)
}
func noempty2(data []string) []string {
	i := 0
	for _, str := range data {
		if str != "" {
			data[i] = str
			i++
		}
	}
	return data[:i]
}
func main05() {
	data := []string{"red", "", "black", "", "", "pink", "blue"}
	afterdata := noempty2(data)
	fmt.Println(afterdata)
}
func Nosame(data []string) []string {
	out := data[:1]
	//遍历原始切片字符串
	for _, word := range data {
		i := 0
		//比较取出的word是否在out中存在
		for ; i < len(out); i++ {
			if word == out[i] {
				break
			}
		}
		if i == len(out) { //表示循环out一圈
			out = append(out, word)
		}
	}
	return out
}
func main06() {
	data := []string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	afterdata := Nosame(data)
	fmt.Println(afterdata)
}
func remove(data []int) []int {
	s := data[3:]
	copy(data[2:], s)
	return data[:4]

}
func main() {
	data := []int{5, 6, 7, 8, 9}
	afterdata := remove(data)
	fmt.Println(afterdata)
}
