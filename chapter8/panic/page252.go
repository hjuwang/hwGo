package main

import "fmt"

func main() {

	PanicDemo4()
}

func PanicDemo4() {

	defer func() {
		a := recover()
		fmt.Println(a)
	}()

	defer fmt.Println("A")

	defer func() {
		fmt.Println("B")
		panic("panic in defer")
		fmt.Println("C")
	}()

	panic("panic")

	fmt.Println("D")
}
