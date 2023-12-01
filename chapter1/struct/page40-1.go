package main

import "fmt"

func main() {

	student := Student{}

	student.UpdateName("张三")

	fmt.Println(student.Name)

}

type Student struct {
	Name string
}

func (s Student) SetName(name string) {

	s.Name = name
}

func (s *Student) UpdateName(name string) {

	s.Name = name

}
