package main

import (
	"errors"
	"fmt"
)

/*
*
判断error 的类型
errorString
wrapError	()
*/
func main() {

	//ExampleBasicError()
	ExampleWrapError()

}

func ExampleBasicError() {

	err := errors.New("this is demo error")

	basicErr := fmt.Errorf("some context: %v", err)

	if _, ok := basicErr.(interface{ Unwrap() error }); !ok {
		fmt.Println("basicErr is errorString")
	}
}

func ExampleWrapError() {

	err := errors.New("this is demo error")

	basicErr := fmt.Errorf("some context: %w", err)

	if _, ok := basicErr.(interface{ Unwrap() error }); ok {
		fmt.Println("basicErr is a wrapError")
	}
}
