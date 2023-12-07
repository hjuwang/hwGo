package main

import "fmt"

func main() {

	//ExampleGreetingWithoutParameter()
	ExampleGreetingWithParameter()
}

func ExampleGreetingWithoutParameter() {
	Greeting("Hello")
}

func ExampleGreetingWithParameter() {

	Greeting("Hi", "Tome", "Jake", "Jim")
}

func Greeting(prefix string, who ...string) {

	if who == nil {
		fmt.Printf("Nobody to say %s", prefix)
		return
	}

	for _, people := range who {
		fmt.Printf("%s %s\n", prefix, people)
	}
}
