

## 原始error
### 普通的error 只是一个接口

```go
type error interface {
    
	Error() string
}
```
### 最基本的实现errorString
```go
type errorString struct {
    //保存error 信息字符串
	s string
}

//实现error interface
func (e *errorString) Error()  {
    return e.string
}
```
### 创建error 
#### 使用errors.New
```go
  //返回一个errorString 对象
  func New(text string) error {
    return &errorString{s}
  }
```  
#### 使用fmt.Errorf() error
只是对errors.New 的简单封装

errors.New 的性能要优于fmt.Errorf


## Go 1.1.3 之后

在fmt 的标准库中新增能加类型 wrapsErrors
```go

type wrapError struct {
msg string
err error
}

func (e *wrapError) Error() string {
return e.msg
}

//这是个新增的方法
func (e *wrapError) Unwrap() error {
return e.err
}
```
msg 保存上下文信息
error 是原始的error

fmt.Errorf 可以根据是否使用%w 格式化动词
返回不同的error 类型

- 如果使用%w 则返回 wrapError
- 如果没使用%w 则返回 errorString
