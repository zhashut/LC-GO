package main

import "fmt"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/1/11
 * Time: 20:45
 * Description: 题意为：[6, 3, 2, 1, 6, 7]，求当前位置数，左边比自己小的数的累加和
 */

func main() {
	arr := []int{6, 3, 2, 1, 6, 7}
	fmt.Println(mergeSort(arr))
}

func mergeSort(arr []int) int {
	if len(arr) == 0 || len(arr) < 2 {
		return 0
	}

	return process(arr, 0, len(arr)-1)
}

// 递归处理
func process(arr []int, l, r int) int {
	if l == r {
		return 0
	}

	mid := l + ((r - l) >> 1)
	return process(arr, l, mid) +
		process(arr, mid+1, r) +
		merge(arr, l, mid, r)
}

// 合并
func merge(arr []int, l, m, r int) int {
	help := make([]int, r-l+1)
	i, p1, p2, res := 0, l, m+1, 0
	for p1 <= m && p2 <= r {
		if arr[p1] < arr[p2] {
			res += (r - p2 + 1) * arr[p1]
			help[i] = arr[p1]
			i, p1 = i+1, p1+1
		} else {
			help[i] = arr[p2]
			i, p2 = i+1, p2+1
		}
	}

	for p1 <= m {
		help[i] = arr[p1]
		i, p1 = i+1, p1+1
	}

	for p2 <= r {
		help[i] = arr[p2]
		i, p2 = i+1, p2+1
	}

	for i := 0; i < len(help); i++ {
		arr[l+i] = help[i]
	}

	return res
}
