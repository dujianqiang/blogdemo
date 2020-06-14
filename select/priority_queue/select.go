//优先队列
package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	c := make(chan int, 100)
	go func() {
		for i:= 0;i<100;i++ {
			a<-i
			b<-i
			c<-i
		}
	}()
	for {
		select {
		case <-a: fmt.Println("recv from a")
		default:
			select {
			case <-b: fmt.Println("recv from b")
			default:
				select {
				case <-c:
					fmt.Println("recv from c")

				default:
					time.Sleep(time.Second)
				}
			}
		}
	}
	time.Sleep(time.Second*100)
}
