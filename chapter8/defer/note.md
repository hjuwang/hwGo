
函数的return 语句并不是原子的，包括两个部分
- 设置返回值
- ret(真正的返回)  
而 函数中的defer 函数栈，就在这其中的两部分中执行

```go

func DeferDemo5() (result int) {

i := 0
defer func() {
result++
}()

return i
}

}
```
此函数的返回值是2

**defer 函数影响主函数的具体名返回值**  

执行的顺序是 :
设置返回值->执行defer函数栈->ret(最终返回)