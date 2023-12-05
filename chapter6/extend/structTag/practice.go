package main

import (
	"fmt"
	"reflect"
)

/*

小练习，改变结构体字段的值
*/

type Fighter struct {
	Name string `json:"name"`
	Age  int    `json:"age"`

	//身体数据
	Height int `json:"height"` //身高
	Weight int `json:"weight"` //体重
	Reach  int `json:"reach"`  //臂展
}

func main() {

	leach := Fighter{
		Name:   "李景亮",
		Age:    34,
		Height: 180,
		Weight: 75,
		Reach:  180,
	}
	fmt.Println(leach.Reach)
	//使用反射修改leach 的臂展为200

	elem := reflect.ValueOf(&leach).Elem()

	elem.FieldByName("Reach").SetInt(200)

	fmt.Println(leach.Reach)
}
