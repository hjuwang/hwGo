package main

import (
	"fmt"
	"reflect"
)

/**
反射第二定律：反射对象和interface 对象之间可以相互转换
*/

func main() {

	var A interface{}

	fmt.Println("A-type:", reflect.TypeOf(A)) //nil
	A = 100

	fmt.Println("A-type:", reflect.TypeOf(A))
	//获取反射对象
	v := reflect.ValueOf(&A)

	//转换成interface{}()s 对象
	i := v.Interface()

	fmt.Println("i-type:", reflect.TypeOf(i))

}
