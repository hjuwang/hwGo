package main

import "fmt"

func main() {

	names := []string{"tome", "jake", "jim"}

	ChangeSliceDemo(names...)

	fmt.Println(names)

}

func ChangeSliceDemo(s ...string) {

	s[0] = "hello"
}
