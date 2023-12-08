package main

import "fmt"

func main() {

	x := make([]int, 0, 10)

	x = append(x, 1, 2, 3)
	y := append(x, 4)
	z := append(x, 5)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
