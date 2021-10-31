// 当一个指针被定义后没有分配到任何变量时，它的值为 nil。
// nil 指针也称为空指针。
// 一个指针变量通常缩写为 ptr。
package main

import "fmt"

func main() {
	var ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr)

	if ptr != nil {
		fmt.Println("不是空指针")
	} /* ptr 不是空指针 */
	if ptr == nil {
		fmt.Println("是空指针")
	} /* ptr 是空指针 */
}
