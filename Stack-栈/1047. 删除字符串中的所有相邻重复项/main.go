package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/2
 * Time: 19:20
 * Description: No Description
 */

func main() {
	s := "abbaca"
	removeDuplicates(s)
}

func removeDuplicates(s string) string {
	var stack []byte

	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
