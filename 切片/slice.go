package main

import "fmt"

/*
	len() 和 cap() 函数
	切片是可索引的，并且可以由 len() 方法获取长度。
*/

func main() {
	var numbers = make([]int, 3, 5)

	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
