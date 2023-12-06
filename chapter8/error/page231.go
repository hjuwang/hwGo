package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	ExampleIs()
}

func ExampleIs() {
	//对os.ErrPermission 包裹两次
	err1 := fmt.Errorf("write file error %w", os.ErrPermission)
	err2 := fmt.Errorf("write file error %w", err1)

	//判断是否含有特定的error 值
	fmt.Println(errors.Is(err2, os.ErrPermission)) //errors.Is(err2,os.ErrPermission)
}
