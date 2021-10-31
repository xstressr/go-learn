/*
	for init; condition; post { }

	init： 一般为赋值表达式，给控制变量赋初值；
	condition： 关系表达式或逻辑表达式，循环控制条件；
	post： 一般为赋值表达式，给控制变量增量或减量。
*/

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
