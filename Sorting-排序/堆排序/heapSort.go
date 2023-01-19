package main

import (
	"fmt"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/1/19
 * Time: 21:20
 * Description: No Description
 */

func main() {
	arr := []int{3, 2, 0, 8, 7, 6, 9, 5}
	heapSort(arr)
	for i := range arr {
		fmt.Printf("%d ", arr[i])
	}
}

// 大根堆
func heapSort(arr []int) {
	if len(arr) == 0 || len(arr) < 2 {
		return
	}

	// O(N*logN)
	/*for i := 0; i < len(arr); i++ {
		headInsert(arr, i)
	}*/

	// O(N)
	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}

	heapSize := len(arr)
	heapSize--
	swap(arr, 0, heapSize)
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		heapSize--
		swap(arr, 0, heapSize)
	}
}

// arr[index]刚来的数，往上
func headInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// arr[index]位置的数，能否往下移动
func heapify(arr []int, index, heapSize int) {
	left := index*2 + 1
	for left < heapSize {
		largest := 0
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		} else {
			largest = left
		}
		if arr[largest] < arr[index] {
			largest = index
		}
		if largest == index {
			break
		}
		swap(arr, largest, index)
		index = largest
		left = index*2 + 1
	}
}

// 交换
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
