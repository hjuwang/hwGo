package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	PrintMethods(time.Time{})
}

func PrintMethods(i interface{}) {

	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i) //也可以使用v.type 获取

	v.NumMethod()
	t.NumMethod()

	for i := 0; i < t.NumMethod(); i++ {

		methodType := t.Method(i).Type
		methodName := t.Method(i).Name
		value := t.Method(i).Func
		fmt.Println("类型:", methodType)
		fmt.Println("方法名:", methodName)
		fmt.Println(value)

		fmt.Println("---------------")
	}

	fmt.Println(v)
	fmt.Println(t)

}
