

#### 常用
向值为nil的map 中添加值会发生panic

delete 值为nil的map不会发生报错

#### 危险操作

map 并不是协程安全的  

map 的操作不是原子的，多个goroutine 同时操作map 有可能会产生读写冲突
