package main

import "fmt"

func main() {

	f := Fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func Fibonacci() func() int {
	a, b := 0, 1

	//闭包
	return func() int {
		a, b = b, a+b
		return a
	}

}
