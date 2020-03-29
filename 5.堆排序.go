package main

import "fmt"

/*
二叉树跟和孩子节点的 索引的关系
0             a[n]


1     a[2*n+1]    a[2*n+2]


*/

// 堆排序的三步
// 1. 构建二叉树
// 2. 构建大顶堆 即每个节点都比左右节点的值大
// 3. 堆顶和 最好一个节点交换，并排除最后一个节点

// 为了和编码对应一般增加一个无用的 0 占据 arr[0] 的位置，或者不用 arr[0]
// 从后往前最后一个节点的父节点角标为 len(arr)/2

func HeapGetMax1(arr []int) []int {
	total := len(arr) - 1

	maxHeap(arr, total) // 左右子节点都小于当前节点   // 从后往前依次调整成大顶堆
	fmt.Println(arr)

	for total > 1 {
		swap(arr, 1, total) //将大的放在数组后面，升序排序
		total--
		heapAdjust(arr, 1, total) // sink 也行  // 替换后 从前往后依次调整
	}

	return arr
}

// 从索引 i 开始调整，到 n，发现不满足大顶堆就调整
// n 当前总节点数
// i 当前索引
func heapAdjust(arr []int, i, n int) {
	for 2*i <= n {
		lchildIndex := 2 * i
		maxChildIndex := lchildIndex
		if n > lchildIndex && arr[lchildIndex+1] > arr[lchildIndex] {
			maxChildIndex = lchildIndex + 1
		}

		if arr[maxChildIndex] > arr[i] {
			arr[i], arr[maxChildIndex] = arr[maxChildIndex], arr[i]
			i = maxChildIndex // 被交换后， 有可能还需要调整，所以用 for
		} else {
			break
		}
	}
}

// 构建大顶堆
func maxHeap(arr []int, total int) {
	for k := total / 2; k >= 1; k-- {
		heapAdjust(arr, k, total)
	}
}

func sink(s []int, k, N int) {
	for {
		i := 2 * k
		if i > N { //保证该节点是非叶子节点
			break
		}
		if i < N && s[i+1] > s[i] { //选择较大的子节点
			i++
		}
		if s[k] >= s[i] { //没下沉到底就构造好堆了
			break
		}
		fmt.Println(k, i)
		swap(s, k, i)
		fmt.Println(k, i)
		k = i
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	arr := []int{0, 1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	// total := len(arr) - 1
	//heapAdjust(arr, total/2, total)
	//sink(arr, total/2, total)
	// maxHeap(arr, total)
	HeapGetMax1(arr)
	fmt.Println(arr)
}
