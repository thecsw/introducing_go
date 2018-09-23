package main

import "fmt"

func fib(x int) int {
	switch x {
	case 0: return 0
	case 1: return 1
	case 2: return 1
	default:return fib(x - 1) + fib(x - 2) 
	}
}

func main() {
	for i:= 0; i < 10; i++ {
		fmt.Println(i, " -- ", fib(i))
	}
}
