package main

import (
	"log"
	"time"
)

func main() {

	TickerDemo()
}

func TickerDemo() {

	ticker := time.NewTicker(1 * time.Second)

	//打印日志
	for c := range ticker.C {
		log.Println("ticker", c)
	}
}
