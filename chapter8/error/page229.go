package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	err := fmt.Errorf("write file error %w", os.ErrPermission)

	deepError := errors.Unwrap(err)

	fmt.Println(deepError)

}
