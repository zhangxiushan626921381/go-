package main

import "fmt"

/*func test(s []int){
	s=append(s,4,5,6)//地址改变但是没给实参传值  所以结果还是[1 2 3]
}

*/
func test(s *[]int) {
	*s = append(*s, 4, 5)
}
func main() {
	s := []int{1, 2, 3}
	test(&s)
	fmt.Println(s)
}
