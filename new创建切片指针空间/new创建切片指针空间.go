package main

import "fmt"

func main() {
	var p *[]int
	p = new([]int)
	fmt.Printf("%p\n", p)
	*p = append(*p, 1, 2)
	for i, d := range *p {
		fmt.Println(i, d)
	}
	fmt.Println(*p)
}
