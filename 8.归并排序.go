package main

import (
	"fmt"
)

func merge(leftarr []int, rightarr []int) []int {
	leftIndex := 0
	rightIndex := 0
	lastArr := []int{}
	for leftIndex < len(leftarr) && rightIndex < len(rightarr) {
		switch {
		case leftarr[leftIndex] < rightarr[rightIndex]:
			lastArr = append(lastArr, leftarr[leftIndex])
			leftIndex++
		case leftarr[leftIndex] > rightarr[rightIndex]:
			lastArr = append(lastArr, rightarr[rightIndex])
			rightIndex++
		case leftarr[leftIndex] == rightarr[rightIndex]:
			lastArr = append(lastArr, rightarr[rightIndex])
			leftIndex++
			rightIndex++
		}
	}
	for leftIndex < len(leftarr) { // 说明有末位有拉下的奇数个, 也应该归并过来
		lastArr = append(lastArr, leftarr[leftIndex])
		leftIndex++
	}
	for rightIndex < len(rightarr) {
		lastArr = append(lastArr, rightarr[rightIndex])
		rightIndex++
	}
	return lastArr
}

// 节省内存
// 不停的分组，最后左右比，按大小排好
func MergeSort(arr []int) []int {
	length := len(arr)
	switch {
	case length <= 1:
		return arr
	case length > 1 && length <= 5:
		return InsertSort(arr)
	default:
		mid := length / 2
		leftarr := MergeSort(arr[:mid])
		rightarr := MergeSort(arr[mid:])

		return merge(leftarr, rightarr)
	}
}

func main() {
	arr := []int{0, 1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(MergeSort(arr))

}
