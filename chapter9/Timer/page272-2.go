package main

import (
	"log"
	"time"
)

/**
延迟某个方法的执行
*/

func main() {

	DelayFunction()
}

func DelayFunction() {
	timer := time.NewTimer(5 * time.Second)

	select {
	case <-timer.C:
		log.Println("Delayed 5s ,start to do something")
	}
}
