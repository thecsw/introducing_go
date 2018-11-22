package main

import ("fmt"
	"time"
)

// chan is bidirectional
func pinger(c chan string) {
	for {
		c <- "ping"
	}
}

// <-chan is sender only
func ponger(c <-chan string) {
	for {
		c <- "pong"
	}
}

// chan<- is receiver only
func printer(c chan<- string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
