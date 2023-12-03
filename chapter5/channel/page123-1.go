package main

import "fmt"

func main() {

	channels := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go Processor(channels[i])
	}

	for i, ch := range channels {
		<-ch
		fmt.Println("Routine", i, "finished")
	}

}

func Processor(ch chan int) {
	//do something

	//可读就不阻塞
	ch <- 0
}
