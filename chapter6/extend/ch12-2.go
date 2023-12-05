package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {

	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x))          // "1"
	fmt.Println(Any(d))          // "1"
	fmt.Println(Any([]int64{x})) // "[]int64 0x8202b87b0"
	fmt.Println(Any([]time.Duration{d}))
}

func Any(i interface{}) string {

	v := reflect.ValueOf(i)
	return formatAtom(v)
}

func formatAtom(v reflect.Value) string {

	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	//基本整形有符号
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	//基本整形无符号
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	// 浮点型
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'g', -1, 64)
	// 布尔类型
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	//
	case reflect.String:
		return strconv.Quote(v.String())
		//其他复杂类型
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}

}
