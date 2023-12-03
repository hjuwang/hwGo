package main

import (
	"context"
	"fmt"
	"time"
)

/*
典型的cancel context 的使用
*/

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go HandRequest(ctx)

	time.Sleep(time.Second * 10)
	fmt.Println("It's time to stop all sub goroutines")
	cancel()

	time.Sleep(time.Second * 5)

}

func HandRequest(ctx context.Context) {
	go WriteRedis(ctx)
	go WriteDatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandRequest done")
			return
		default:
			fmt.Println("HandRequest running")
			time.Sleep(time.Second * 2)
		}
	}

}

func WriteRedis(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Write Redis done")
			return
		default:
			fmt.Println("Write Redis running")
			time.Sleep(time.Second * 2)
		}
	}
}

func WriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Write Database done")
			return
		default:
			fmt.Println("Write Database running")
			time.Sleep(time.Second * 2)
		}
	}
}
