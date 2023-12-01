package main

import "fmt"

func SilceCap() {

	var array [10]int
	var slice = array[5:6]

	fmt.Printf("len(slice) = %d\ncap(slice) = %d\n", len(slice), cap(slice))
}

func main() {

	SilceCap()
}
