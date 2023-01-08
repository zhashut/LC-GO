package main

import "fmt"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/1/8
 * Time: 10:54
 * Description: No Description
 */

func main() {
	arr := []int{3, 2, 5, 4, 6, 7, 3}
	sort := insertSort(arr)
	fmt.Println(sort)
}

func insertSort(arr []int) []int {
	if len(arr) == 0 || len(arr) < 2 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}

	return arr
}
