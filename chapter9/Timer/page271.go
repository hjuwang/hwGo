package main

/**
应用场景，设定超时时间
*/
import (
	"fmt"
	"time"
)

func main() {

	WaitChannel(make(chan string))
}

func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second)

	for {
		select {
		case <-conn:
			timer.Stop()
			return true
		case <-timer.C: //定时结束（定时器时间到）
			fmt.Println("WaitChannel time timeout")
			return false

		}
	}

}
