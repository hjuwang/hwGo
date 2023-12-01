package main

import (
	"runtime"
	"time"
)

func main() {
	//定义一个P
	runtime.GOMAXPROCS(1)

	go func() {
		for {
			//无函数调用的无线循环
		}
	}()

	time.Sleep(time.Second)

	println("Done")
}
