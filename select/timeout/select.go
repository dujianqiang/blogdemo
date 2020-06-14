package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make (chan bool, 1)
	go func() {
		time.Sleep(time.Second*10) // sleep one second
		timeout <- true
	}()
	select {
	case <- timeout:
		fmt.Println("timeout!")
	}
}
