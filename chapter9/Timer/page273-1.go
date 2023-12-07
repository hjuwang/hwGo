package main

/*
匿名定时器的使用
*/
import (
	"log"
	"time"
)

func main() {
	AfterDemo()
}

func AfterDemo() {

	log.Println(time.Now())

	<-time.After(1 * time.Second)
	log.Println(time.Now())
}
