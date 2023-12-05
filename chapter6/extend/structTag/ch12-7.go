package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

/**
将req.请求解析到结构体字段

*/

func main() {

}

func search() {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10
}

func Unpack(req *http.Request, ptr interface{}) error {

	if err := req.ParseForm(); err != nil {
		return err
	}
	// 用于保存结构体字段的值
	fields := make(map[string]reflect.Value)
	////ptr 其实是指针变量

	v := reflect.ValueOf(ptr).Elem()

	for i := 0; i < v.NumField(); i++ {

		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag

		//根据标签名称获取标签内容
		name := tag.Get("http")

		if name == "" {
			//转换成结构体字段名的小写
			name = strings.ToLower(fieldInfo.Name)
		}

		fields[name] = v.Field(i)
	}

	//更新结构体
	for name, values := range req.Form {

		//f 始终表示结构体的字段
		f := fields[name]
		//如果f 是零值就continue
		if !f.IsValid() {
			continue
		}

		for _, value := range values {
			if f.Kind() == reflect.Slice {
				//创建一个指向该类型的指针
				fieldptr := reflect.New(f.Type().Elem())
				//获取可以寻址的指针值
				elem := fieldptr.Elem()

				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s:%v", name, err)
				}

			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s:%v", name, err)
				}
			}
		}

	}

	return nil
}

func populate(v reflect.Value, value string) error {

	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		//解析10进制数据
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Bool:
		parseBool, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(parseBool)
	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
