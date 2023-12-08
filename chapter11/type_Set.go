package main

import "fmt"

/*
使用泛型实现一个Set 集合

Set 可以存储一组不重复的数据
*/
func main() {
	//测试这个新实现的set

	set := make(Set[int])
	set.Add(1)
	set.Add(2)
	fmt.Println(set.Contains(3)) //false
	fmt.Println(set.Contains(2)) //true
	fmt.Println(set.Len())       //2
	set.Delete(2)

	fmt.Println(set.Contains(2)) //false

}

type Set[T comparable] map[T]struct{}

func MakeSet[T comparable]() Set[T] {
	return make(map[T]struct{})
}

// Add 向set 中添加元素
func (s Set[T]) Add(v T) {

	s[v] = struct{}{}
}

//Delete 在Set 中删除元素

func (s Set[T]) Delete(v T) {

	delete(s, v)

}

//Contains 判断Set 中是否存在元素 v

func (s Set[T]) Contains(v T) bool {

	_, ok := s[v]
	return ok
}

//返回Set 的大小

func (s Set[T]) Len() int {

	return len(s)
}

//遍历 逐个调用函数f

func (s Set[T]) iterate(f func(T)) {
	for k := range s {
		f(k)
	}
}
