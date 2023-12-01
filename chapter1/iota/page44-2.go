package main

import "fmt"

const (
	Locked = 1 << iota
	Woken
	Starving
	WaiterShift           = iota
	StarvationThresholdNs = 1e6 //1e6 表示1*10^6
)

func main() {

	fmt.Println(Locked)
	fmt.Println(Woken)
	fmt.Println(Starving)
	fmt.Println(WaiterShift)
	fmt.Println(StarvationThresholdNs)
}
