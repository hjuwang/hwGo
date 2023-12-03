package main

import (
	"context"
	"time"
)

func main() {

	ctx, cancel2 := context.WithCancel(context.Background())

	cancel2()
	deadline, cancelFunc := context.WithDeadline(context.Background(), time.Now())
	timeout, c := context.WithTimeout(context.Background(), time.Second)
	value := context.WithValue(context.Background(), "key", "value")

}
