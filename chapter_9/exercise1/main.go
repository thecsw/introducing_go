package main

import ("fmt"
	m "github.com/introducing_go/chapter_9/exercise1/math"
)
	
func main() {
	xs := []float64{1,2,3,4}
	avg := m.Average(xs)
	fmt.Println(avg)
}
