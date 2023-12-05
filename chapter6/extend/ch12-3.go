package main

import (
	"fmt"
	"reflect"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func main() {

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	Display("strangelove", reflect.ValueOf(strangelove))
	//Display("strangelove", strangelove)
}

func Display(path string, v reflect.Value) {

	switch v.Kind() {
	//空的 reflect.Value 的 kind 即为 Invalid(表示空值)
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	//切片和数组
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {

			//v.Index(i)Index(i)获得索引i对应的元素
			Display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))

		}
	//结构体
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			//v.Field(i)Field(i)获得索引i对应的字段
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			Display(fieldPath, v.Field(i))
		}
		//map
	case reflect.Map:
		for _, key := range v.MapKeys() {
			mapth := fmt.Sprintf("%s[%v]", path, key)
			//MapIndex(key)返回map中key对应的value
			Display(mapth, v.MapIndex(key))

		}
		//指针
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			//Elem()返回指针指向的值
			Display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
		//接口
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			Display(path+".value", v.Elem())
		}

	default:
		fmt.Printf("%s = %v\n", path, v)
	}

}
