// 不定参函数
package main

import (
	"fmt"
)

func add(arr ...int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	fmt.Println(sum)

}
func plus(are ...int) {
	sum := 0
	for _, a := range are {
		sum = sum + a
	}
	fmt.Println(sum)
}
func main() {
	plus(1, 2, 3, 455)
	add(1, 2, 1, 32, 656)
}
