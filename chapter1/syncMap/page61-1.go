package main

import (
	"fmt"
	"sync"
)

var m sync.Map

func main() {

	//增加（或修改）
	m.Store("Jim", 80)
	m.Store("Kevin", 90)
	m.Store("Jim", 90)

	//查询
	score, _ := m.Load("Jim")
	fmt.Printf("Jim's score is %d\n", score.(int))

	m.Range(func(key, value interface{}) bool {
		fmt.Printf("key is %v, value is %v\n", key, value)
		return true
	})

}
