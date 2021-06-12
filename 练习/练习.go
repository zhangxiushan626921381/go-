package main

import "fmt"

func main() {
	//
	var arr [20]byte
	for i := 0; i < len(arr); i++ {
		fmt.Scanf("%c", &arr[i])
	}
	m := make(map[byte]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	for k, v := range m {
		if v > 0 {
			fmt.Printf("%c %d\n", k, v)
		}
	}
}
