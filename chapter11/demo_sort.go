package main

import (
	"fmt"
	"sort"
)

/*
实现一个可排序的泛型接口
*/

// Order 定义类型集合接口
type Order interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | //有符号
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | //无符号
		~float32 | ~float64 | //浮点
		~string //字符串
}

// OrderSlice 泛型类型
type OrderSlice[T Order] []T

func (o OrderSlice[T]) Len() int {
	return len(o)
}

func (o OrderSlice[T]) Less(i, j int) bool {
	return o[i] < o[j]
}

func (o OrderSlice[T]) Swap(i, j int) {

	//交换
	o[i], o[j] = o[j], o[i]
}

func main() {

	//排序字符串 OrderSlice[byte] 可以存放byte切片
	byteStr := OrderSlice[byte]("sdfadsfa")
	sort.Sort(byteStr)
	fmt.Println(byteStr)
	fmt.Println(string(byteStr))

	//排序整形数字

	nums := OrderSlice[int]{1, 4, 7, 10, 34, 100, 2}
	sort.Sort(nums)
	fmt.Println(nums)

	//排序浮点数字

	floatNums := OrderSlice[float64]{1.0, 2.3, 0.7, 0.01, 0.001, 2.89}
	sort.Sort(floatNums)
	fmt.Println(floatNums)
}
