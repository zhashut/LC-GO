package main

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/20
 * Time: 21:29
 * Description: No Description
 */

func main() {
	s := "aab"
	partition(s)
}

func partition(s string) [][]string {
	result := make([][]string, 0)

	backtrack(s, &result, []string{}, 0)
	return result
}

func backtrack(s string, result *[][]string, path []string, startIndex int) {
	if startIndex >= len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	for i := startIndex; i < len(s); i++ {
		if !isPalindrome(s, startIndex, i) {
			continue
		}
		path = append(path, s[startIndex:i+1])
		backtrack(s, result, path, i+1)
		path = path[:len(path)-1]
	}
}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
