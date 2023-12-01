package main

import (
	"fmt"
	"time"
)

func addNumberToChan(c chan int) {

	for {
		c <- 1
		time.Sleep(time.Second * 1)
	}
}

func main() {

	var chan1 = make(chan int, 10)
	var chan2 = make(chan int, 10)

	//并发写数据
	go addNumberToChan(chan1)
	go addNumberToChan(chan2)

	//select 监听 channel
	for {

		select {
		case e := <-chan1:
			fmt.Printf("Get element from chan1: %d\n", e)
		case e := <-chan2:
			fmt.Printf("Get element from chan2: %d\n", e)
		default:
			fmt.Print("No element in chan1 and chan2\n")
			time.Sleep(time.Second * 1)
		}
	}

}
