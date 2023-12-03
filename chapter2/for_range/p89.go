package main

func main() {

}

func RangeChannel() {

	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	for e := range c {
		println(e)
	}
}
