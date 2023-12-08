package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {

		fmt.Printf("内存地址:%v,i=%v\n", &i, i)
	}
}
