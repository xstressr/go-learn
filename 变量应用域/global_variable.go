package main

import "fmt"

// 声明全局变量
var g int

func main() {
	// 声明局部变量
	var a, b int

	// 初始化参数
	a = 10
	b = 20
	g = a + b
	fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)
}

// 加入全局变量和局部变量是一个名字，默认使用局部变量
