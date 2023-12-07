package main

import (
	"log"
	"time"
)

/**
异步延迟调用
*/

func main() {

	AfterFuncDemo()
}

func AfterFuncDemo() {
	log.Println("AfterFuncDemo start:", time.Now())

	//这是异步执行
	time.AfterFunc(1*time.Second, func() {
		log.Println("AfterFuncDemo end:", time.Now())
	})

	//因为time.AfterFunc是异步执行，所以这里需要主要携程等待

	time.Sleep(2 * time.Second)
}
