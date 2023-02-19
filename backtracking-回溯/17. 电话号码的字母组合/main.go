package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/19
 * Time: 21:19
 * Description: No Description
 */

var (
	m    []string
	path []byte
	res  []string
)

func main() {
	digits := "23"
	letterCombinations(digits)
}

func letterCombinations(digits string) []string {
	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path, res = make([]byte, 0), make([]string, 0)
	if digits == "" {
		return res
	}
	dfs(digits, 0)
	return res
}

func dfs(digits string, index int) {
	if len(path) == len(digits) {
		temp := string(path)
		res = append(res, temp)
		return
	}

	digit := int(digits[index] - '0')
	str := m[digit-2]
	for i := 0; i < len(str); i++ {
		path = append(path, str[i])
		dfs(digits, index+1)
		path = path[:len(path)-1]
	}
}
