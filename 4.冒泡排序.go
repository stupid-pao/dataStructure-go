package main

import "fmt"

func BubbleSort(arr []int) []int {
	length := len(arr)
	if length < 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			for j := 0; j < length-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}
	return arr
}

func main4() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(BubbleSort(arr))

}
