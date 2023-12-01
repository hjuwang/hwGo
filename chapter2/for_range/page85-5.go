package main

import "fmt"

func main() {
	RangeDemo()
}

func RangeDemo() {
	s := "你好萝卜"
	for i, v := range s {
		fmt.Println(i, v)
	}
}
