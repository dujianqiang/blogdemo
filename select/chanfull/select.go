package main

import "fmt"

func main() {
	ch2 := make(chan int, 1)
	ch2 <- 1
	select {
	case ch2 <- 2:
	default:
		fmt.Println("channel is full !")
	}
}
