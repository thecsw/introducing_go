package main

import "fmt"

func half(x int) (half int, even bool) {
	half = x / 2
	if x % 2 == 0 {
		even = true
	} else {
		even = false
	}
	return
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}
