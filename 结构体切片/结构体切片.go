package main

import "fmt"

type student struct {
	id   int
	name string
	age  int
	sex  string
}

func main() {

	var stu []student = make([]student, 3)
	p := &stu //结构体切片指针
	(*p)[0] = student{102, "dd", 111, "sss"}
	(*p)[1] = student{102, "aa", 111, "sss"}
	(*p)[2] = student{102, "ff", 111, "sss"}
	*p = append(*p, student{102, "ss", 111, "sss"})
	fmt.Printf("%T\n", p)
	fmt.Println(*p)
	for i := 0; i < len(*p); i++ {
		fmt.Println((*p)[i])
	}

}
