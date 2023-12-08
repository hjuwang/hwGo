### 函数泛型

```go
func SumIntsOrFloats[K comparable, V int | float64](m map[K]V) V {
    
}
```

在函数名后使用[] 来定义泛型类型 K,V  
`[K comparable, V int | float64]`  
K comparable 表示类型 必须是可比较的类型  
V int | float64 表示类型 V 可以是 int 类型或者 float64 类型

### 类型泛型

-  使用接口定义类型集合  
   ```go
        type Order interface {
            ~int | ~int8 | ~int16 | ~int32 | ~int64 |
                ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
                ~float32 | ~float64 |
                ~string
        }
   ```
   **注意:**
    - `~<type>` 来指定一组类型，只要底层类型为type 就包含在这个集合中
    - 波浪号"~"后的的类型为底层类型
    - 类型取交集,定义每一行为子集,结果取这些子集的交集,上述例子是一行
- 定义类型泛型
    ```go
     type OrderSlice[T Order] []T
    ```
    **注意：**  
    - 定义的泛型类型必须初始化以后才可以使用
      ```go
       eStr := OrderSlice[byte]("sdfadsfa")
      ```

