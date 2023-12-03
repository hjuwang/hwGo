package main

import (
	"context"
	"fmt"
	"time"
)

/*
典型的timer context 的使用
*/

func main() {

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	go HandRequest(ctx)

	//cancel()
	time.Sleep(time.Second * 11)

	fmt.Println(ctx.Err())

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
