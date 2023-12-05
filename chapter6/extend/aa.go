package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {

	var w io.Writer = os.Stdout

	fmt.Println(w, reflect.TypeOf(w))
	/*
		因为reflect.TypeOf(w) 返回的是一个动态类型的接口值
		又因为w的动态类型是 *os.File，所以打印的是*os.File
	*/

}
