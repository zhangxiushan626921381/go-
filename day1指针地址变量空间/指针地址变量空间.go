package main

import "fmt"

func main01() {
	var a int = 10
	var p *int = &a
	*p = 200 //借助a变量的地址，操作a对应空间
	b := *p
	fmt.Println(b)
	//在heap上申请一片内存空间
	p = new(int) //默认类型的默认值
	var k *string
	k = new(string)
	fmt.Printf("%q\n", *k) //打印Go语言格式的字符串
}

/*func main()  {
	var p *int=0  //err  野指针
	fmt.Println(*p)
}*/
func swap(a int, b int) {
	a, b = b, a
	fmt.Println("a=", a, "b=", b)
}
func swap2(x, y *int) {
	*x, *y = *y, *x
	fmt.Println("a=", *x, "b=", *y)
}
func main() {
	a, b := 10, 20
	swap(a, b) //swap执行完释放
	fmt.Println("a=", a, "b=", b)
	swap2(&a, &b) //传地址值
	fmt.Println("a=", a, "b=", b)
}
