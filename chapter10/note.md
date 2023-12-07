

### ":=" 语法糖

在相同的作用域内，使用":=" 重新声名同名便变量，不会产生新的变量
### 函数的可变参数

- 可变参数必须在函数列表的尾部
- **可变参数在函数内部是作为切片来解析的**
- **可变参数可以不填，不填时函数内部当成nil 切片来处理**
- 可变参数必须是相同的类型

注意：
- 切片传入函数的可变参数的时候不会产生新的切片
- 函数内部使用的切片与传入的切片共享存储空间