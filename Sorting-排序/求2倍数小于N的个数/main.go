package main

import "fmt"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/1/11
 * Time: 22:36
 * Description: 题意为：当前下标数，右边的数的2倍比它小的的话，ans += 1， 遍历整个数组，就累计的个数
 */

func main() {
	arr := []int{1, 4, 6, 8, 9, 1, 1, 2, 3, 3, 4}
	fmt.Println(reversePairs(arr))
}

func reversePairs(arr []int) int {
	if len(arr) == 0 || len(arr) < 2 {
		return 0
	}

	return process(arr, 0, len(arr)-1)
}

func process(arr []int, l, r int) int {
	if l == r {
		return 0
	}

	mid := l + ((r - l) >> 1)
	return process(arr, l, mid) + process(arr, mid+1, r) + merge(arr, l, mid, r)
}

func merge(arr []int, l, m, r int) int {
	ans := 0
	windowR := m + 1
	for i := l; i <= m; i++ {
		for windowR <= r && arr[i] > (arr[windowR]<<1) {
			windowR++
		}
		ans += windowR - m - 1
	}

	help := make([]int, r-l+1)
	i, p1, p2 := 0, l, m+1
	for p1 <= m && p2 <= r {
		if arr[p1] < arr[p2] {
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

	for i = 0; i < len(help); i++ {
		arr[l+i] = help[i]
	}

	return ans
}
