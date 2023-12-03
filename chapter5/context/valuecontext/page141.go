package main

import (
	"context"
	"fmt"
	"time"
)

/*
value context 的用法
*/

func main() {
	ctx := context.WithValue(context.Background(), "parameter", "1")
	go HandlerRequest(ctx)

	time.Sleep(time.Second * 5)
}

func HandlerRequest(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandlerRequest done")
			return
		default:
			fmt.Println("HandlerRequest running, parameter is ", ctx.Value("parameter"))
			time.Sleep(time.Second * 2)
		}
	}

}
