package main

import "fmt"

func main()  {
	var num []int = make([]int, 3, 5)
	var num1 = make([]int, 3, 5)
	fmt.Println(len(num))
	fmt.Println(cap(num))
	fmt.Println(num1)

	numbers := []int{0,1,2,3,4,5,6,7,8}
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	fmt.Println("numbers[1:4] ==", numbers[:4])

}
