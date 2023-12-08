package main

func main() {

	SumIntsOrFloats(map[string]float64{"a": 1.1, "b": 2.1})

	//这种类型需要初始化才能使用
	type NewV[T any] []T

}

func SumIntsOrFloats[K comparable, V int | float64](m map[K]V) V {

	var sum V
	for _, v := range m {
		sum += v
	}

	return sum
}
