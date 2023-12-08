package benchmark

func MakeSliceWithoutAlloc() []int {

	var newSlice []int

	for i := 0; i < 100000; i++ {
		newSlice = append(newSlice, i)
	}

	return newSlice
}

func MakeSliceWithAlloc() []int {

	var newSlice []int

	//预先分配内存空间
	newSlice = make([]int, 0, 100000)

	for i := 0; i < 100000; i++ {
		newSlice = append(newSlice, i)
	}

	return newSlice
}
