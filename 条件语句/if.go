package main

import "fmt"

func main() {
	/* 定义全局变量 */
	var a int = 10

	/* 使用if语句判断布尔表达式 */
	if a < 20 {
		/* 如果条件为true则执行一下语句 */
		fmt.Printf("a 小于 20 \n")
	}
	fmt.Printf("a 的值为： %d\n", a)
}
