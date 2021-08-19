package main

import (
	"fmt"
	"strings"
)

func main() {
	b := "G900#kk#r#"
	r := strings.NewReplacer("#", "*")
	f := r.Replace(b)
	fmt.Println(f)
}
