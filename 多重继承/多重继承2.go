package main

import "fmt"

type D struct {
	name string
}
type E struct {
	age int
}
type F struct {
	D
	E
	id int
}

func main() {
	var s F
	s.age = 15
	s.id = 10
	s.name = "姆巴佩"
	fmt.Println(s)
	var sd F = F{D{"梅西"}, E{18}, 23}
	fmt.Println(sd)
}
