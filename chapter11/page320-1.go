package main

import "fmt"

func main() {

	keys := MapKeys(map[string]float64{"a": 1.1, "b": 2.1})

	fmt.Println(keys)

}

func MapKeys[K comparable, V any](m map[K]V) []K {

	//实现定义底层数组容量 len(m)
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
