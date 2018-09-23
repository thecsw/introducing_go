package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		var word string
		if i % 3 == 0 {
			word += "Fizz"
		}
		if i % 5 == 0 {
			word += "Buzz"
		}
		if len(word) != 0 {
			fmt.Println(word)
		} else {
			fmt.Println(i)
		}
	}
}
