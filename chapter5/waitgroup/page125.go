package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		//Do some work
		time.Sleep(time.Second)

		fmt.Println("routine 1 finished")
		wg.Done()

	}()

	go func() {
		//Do some work
		time.Sleep(time.Second)

		fmt.Println("routine 2 finished")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("All routines finished")
}
