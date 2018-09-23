package main

import "fmt"

func find_max(args ...int) (max int) {
	max = args[0]
	for _, value := range args {
		if value > max {
			max = value
		}
	}
	return
}

func main() {
	x := []int{1 , 3 , 32 , 32 , 93, 0 , 21 }
	fmt.Println(find_max(1, 2, 3, 4, 99))
	// If we want to pass a slice, it needs to be followed
	// with an ellipsis
	fmt.Println(find_max(x...))
}
