package main

/**
反面实例,recover 没有直接在defer 函数中调用,导致recover 无效
*/
import "fmt"

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
