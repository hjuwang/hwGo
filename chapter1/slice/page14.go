package main

import (
	"fmt"
)

func main() {
	SlicePrint()
}

func SliceRise(s []int) {

	fmt.Printf("%p\n", s)

	s = append(s, 0)

	fmt.Printf("%p\n", s)

	for i := range s {
		s[i]++
	}

}

func SlicePrint() {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)

	//fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)

	//SliceRise(s1)
	//fmt.Println(s1)

	SliceRise(s2)
	fmt.Println(s2)

}
