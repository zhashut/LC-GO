package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/28
 * Time: 22:37
 * Description: No Description
 */

func main() {

}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '(') || (v == '[') || (v == '{') {
			stack = append(stack, v)
		} else if ((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
