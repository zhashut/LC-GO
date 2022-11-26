package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/27
 * Time: 0:10
 * Description: No Description
 */

func main() {
	s := "abcdefg"
	reverseLeftWords(s, 2)
}

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	return string(b[n:]) + string(b[:n])
}

func reverseLeftWords1(s string, n int) string {
	b := []byte(s)
	result := ""

	for i := 0; i < len(s); i++ {
		if i == n {
			result += string(b[i:])
			result += string(b[:i])
			break
		}
	}
	return result
}
