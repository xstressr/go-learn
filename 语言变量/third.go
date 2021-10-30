package main

import "fmt"

func main() {
	/*
		声明变量的一般形式是使用 var 关键字：
	*/
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	/* var intVal int
	intVal :=1  */
	// 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明

	intVal := 1
	fmt.Print(intVal)

	// 使用 := 初始化省略变量的类型有系统自动判断
}
