

#### 作用于 array string slice


返回下标和元素的值，如果只有一个元素则返回**下标**

#### 作用于channel
返回channel 中的数据

range 会阻塞等待channel中的数据
如果 range 值为nil 则会永久阻塞


**for range 表达式在遍历开始前就已经决定了循环的次数** (对于array 和 slice)