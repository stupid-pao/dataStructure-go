package main

import "fmt"

// 快速排序排的最快
func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		splitdata := arr[0] // 以第一個为基准
		low := make([]int, 0, 0)
		high := make([]int, 0, 0)
		mid := make([]int, 0, 0)

		mid = append(mid, splitdata)
		for i := 1; i < length; i++ {
			switch {
			case arr[i] < splitdata:
				low = append(low, arr[i])
			case arr[i] > splitdata:
				high = append(high, arr[i])
			case arr[i] == splitdata:
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high)
		myarr := append(append(low, mid...), high...)
		return myarr
	}

}

func main() {
	arr := []int{0, 1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(QuickSort(arr))

}
