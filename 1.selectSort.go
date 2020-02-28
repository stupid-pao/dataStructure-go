package main

import "fmt"

// 选择排序

func SelectSortMax(arr []int) int {
	length := len(arr)
	if length < 1 {
		return arr[0]
	}
	max := arr[0]
	for i := 1; i < length; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func SortMax(arr []int) []int {
	length := len(arr)
	if length < 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			maxIndex := i
			for j := i; j < length; j++ {
				if arr[j] > arr[maxIndex] {
					maxIndex = j
				}
			}
			if i != maxIndex {
				arr[maxIndex], arr[i] = arr[i], arr[maxIndex]
			}
		}
	}
	return arr
}

func main1() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0}
	fmt.Println(SortMax(arr))

}
