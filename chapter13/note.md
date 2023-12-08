

## 常见的编程陷阱
### Slice扩容
- Slice 扩容可能会产生新的切片,注意这个**新** 字
- 使用append()函数的时候,append可能会追加nil值,应该尽量避免追加无意义的元素
    ```go
  //不会发生编译报错
    slice = append(slice,nil)
    ```
- append() 不会修改原切片的值,切片的实现原理是数组array

### 循环变量引用
- 在Go中 for 循环变量i 只分配一次地址
  ```go

  func main() {
    for i := 0; i < 3; i++ {
		    fmt.Printf("内存地址:%v,i=%v\n", &i, i)
	}
  }
  ```
  打印:
  ```
  内存地址:0xc00000a0b8,i=0
  内存地址:0xc00000a0b8,i=1
  内存地址:0xc00000a0b8,i=2
  ```
- 启动协程使用循环变量   
       如果循环体中启动协程,且协程会使用循环变量,
       则协程可能会在循环结束后才启动,此时循环使用的循环变量了能已经被改写
      例如:
     ```go
        func Process1(tasks []string){
			    for _,task := range tasks{
                    go func (){
                        fmt.Pringf("Work start Process task: %s\n",task)
                    }()         
                }      
        }

     ```


### recover() 失效
- recover() 函数必须在defer() 函数中直接调用,否则会失效  
失效的例子
   ```go
    package main

    /**
    反面实例,recover 没有直接在defer 函数中调用,导致recover 无效
    */

    func IsPanic() bool {
    
        if err := recover(); err != nil {
            fmt.Println("Recover success:", err)
            return true
        }
    
        return false
    }

    func UpdateTable() {
    
        //根据是否发生panic 判断是否提交数据库事务
        defer func() {
            if IsPanic() {
                fmt.Println("UpdateTable panic")
                //更新表操作回滚
            } else {
                //提交表更新
            }
        }()
    
        //数据库表更新操作
    }


   ```
- 另外两个recover() 失效的条件
  - panic() 的参数为nil
  - 当前协程没有发生panic(这不就是没用到吗)
