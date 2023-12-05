package main

import (
	"fmt"
	"reflect"
)

/*
*
反射第一定律：反射可以将interface 类型的变量转换为 反射对象
*/
func main() {

	var x float64 = 3.4
	//这里注意：reflect.TypeOf(i any) 接受的是any 类型的参数
	t := reflect.TypeOf(x)

	fmt.Println("type:", t)

	v := reflect.ValueOf(x)
	fmt.Println("value:", v)

}
