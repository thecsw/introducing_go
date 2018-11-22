package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
	fmt.Println(&x, &y)
	fmt.Printf("%#v, %T\n", x, y)
	
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("AAAAAAAARGGHHH")
}
