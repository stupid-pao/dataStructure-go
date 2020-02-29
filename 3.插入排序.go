package main

import "fmt"

func InsertSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 1; i < length; i++ { // 跳过第一个
			backup := arr[i] // 备份要插入的数据
			j := i - 1       // 上一个位置 开始向前循环插入
			for j > 0 && backup < arr[j] {
				arr[j+1] = arr[j] // 小就依次向后挪动一位
				j--
			}
			arr[j+1] = backup // 插入备份值
		}
		return arr
	}
}

func main3() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(InsertSort(arr))

}
