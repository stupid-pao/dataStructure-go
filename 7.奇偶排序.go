package main

import (
	"fmt"
)

// 奇数位 和 偶数位开始,交替进行 冒泡
// 可以多线程跑？
//
func OddEven(arr []int) []int {
	isSorted := false // 奇偶数位都不需要排序的话

	for isSorted == false {
		isSorted = true
		for i := 1; i < len(arr)-1; i += 2 {
			if arr[i] < arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}

		}
		for i := 0; i < len(arr)-1; i += 2 {
			if arr[i] < arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}

		}
	}
	return arr
}

func main() {
	arr := []int{0, 1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(OddEven(arr))

}
