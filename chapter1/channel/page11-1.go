package main

func readChan(c <-chan int) {
	<-c
	//有缓冲，无数据，没关闭，读数据，会阻塞
}

func writeChan(c chan<- int) {
	c <- 1
}

func main() {

	var mychan = make(chan int, 10)

	readChan(mychan)
	writeChan(mychan)
}
