
### 接口基本概念

Interface 变量可以存储任意**实现**了该接口类型的变量

### 反射三定律

- 反射可以将interface类型变量转换成反射对象
- 反射可以将反射对象还原成interface对象
- 反射对象可修改，value值必须是可设置的

### Go圣经学习

interface的结构类似如下
```go
interface{
	type  //接口的动态类型
	value //接口的动态值
}

```

reflect.TypeOf(i interface{}) 返回的是接口i的动态类型  
reflect.ValueOf(i interface{}) 返回的是接口i 的动态值
