

空select语句会引起永久阻塞
例如：

```go
package main

func main() {
	select {}
	
}

```
输出
```powersehll
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [select (no cases)]:

```