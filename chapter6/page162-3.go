package main

import (
	"fmt"
	"reflect"
)

/*
反射第三定律：反射对象可以修改，value值必须是可设置的
*/
func main() {

	var x float64 = 3.4

	v := reflect.ValueOf(&x)

	v.Elem().SetFloat(5.6)

	fmt.Println(x)
}
