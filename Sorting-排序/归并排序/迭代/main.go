package main

import "fmt"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/1/11
 * Time: 0:12
 * Description: No Description
 */

func main() {
	arr := []int{3, 2, 4, 1, 5, 2, 9, 8}
	mergeSort(arr)
	for i := range arr {
		fmt.Print(arr[i])
	}
}

func mergeSort(arr []int) {
	if len(arr) == 0 || len(arr) < 2 {
		return
	}
	n := len(arr)
	// 步长
	mergeSize := 1
	for mergeSize < n {
		l := 0
		for l < n {
			if mergeSize >= n-l {
				break
			}
			m := l + mergeSize - 1
			r := m + min(mergeSize, n-m-1)
			merge(arr, l, m, r)
			l = r + 1
		}
		// 防止溢出
		if mergeSize > n/2 {
			break
		}
		mergeSize <<= 1
	}
}

func merge(arr []int, L, M, R int) {
	help := make([]int, R-L+1)
	i, p1, p2 := 0, L, M+1
	for p1 <= M && p2 <= R {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			i, p1 = i+1, p1+1
		} else {
			help[i] = arr[p2]
			i, p2 = i+1, p2+1
		}
	}

	for p1 <= M {
		help[i] = arr[p1]
		i, p1 = i+1, p1+1
	}

	for p2 <= R {
		help[i] = arr[p2]
		i, p2 = i+1, p2+1
	}

	for i = 0; i < len(help); i++ {
		arr[L+i] = help[i]
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
