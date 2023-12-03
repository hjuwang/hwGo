

valueCtx 在父节点context 中设置的k-v 可以在子节点携程中获取
参考page141.go


### cancelCtx
通过 context.withCancel() 返回context 和cancelFunc 方法  
调用cancelFunc() 方法可以结束整个goroutine树d  

### timerCtx
通过 context.withDeadline(time,time.Time) 返回context,cancelFunc  
时间到time 自动结束goroutine 树，在时间到time 之前可以使用cancelFunc() 强制结束goroutine树  
自动结束 Err() 返回 `context deadline exceeded`  
使用cancelFunc() 强制结束返回 `context canceled` 
