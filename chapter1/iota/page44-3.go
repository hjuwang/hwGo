package main

import "fmt"

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1
	bit1, mask1
	_, _
	bit3, mask3
)

func main() {

	fmt.Println(bit0, mask0)
	fmt.Println(bit1, mask1)
	fmt.Println(bit3, mask3)

}
