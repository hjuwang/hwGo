package main

import "fmt"

func main() {

	fmt.Println(DeferDemo5())
}

func DeferDemo5() (result int) {

	i := 0
	defer func() {
		result++
	}()

	return i
}
