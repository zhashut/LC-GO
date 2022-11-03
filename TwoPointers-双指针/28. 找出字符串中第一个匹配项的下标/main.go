package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/3
 * Time: 0:23
 * Description: No Description
 */

func main() {
	haystack, needle := "hello", "ll"
	strStr(haystack, needle)
}

func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}
