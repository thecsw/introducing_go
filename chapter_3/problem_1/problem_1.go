package main

import "fmt"

func main() {
	fmt.Print("Enter Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	celcius := (input - 32) * 5 / 9
	fmt.Println("Celcius: ", celcius)
}
