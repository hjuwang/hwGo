package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	ExampleAssetChanWithAs()

}

func ExampleAssetChanWithAs() {

	err := os.PathError{
		Op:   "open",
		Path: "a.txt",
		Err:  os.ErrPermission,
	}

	err2 := fmt.Errorf("some context: %w", err)

	var target *os.PathError

	fmt.Println(errors.As(err2, &target))
}
