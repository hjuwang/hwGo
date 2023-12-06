package main

import (
	"reflect"
)

func main() {

	a := "aaa"

	t := reflect.ValueOf(&a).Type()
	//elem := t.Elem()
	elem := t.Elem()

	//fmt.Println(t)
	//fmt.Println(a)
	//fmt.Println(elem)

}
