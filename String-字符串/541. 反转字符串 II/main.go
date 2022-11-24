package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/24
 * Time: 23:17
 * Description: No Description
 */

func main() {
	s := "abcdefg"
	k := 2

	reverseStr(s, k)
}

func reverseStr(s string, k int) string {
	if k > len(s) {
		k = len(s)
	}
	for i := 0; i < len(s); i = i + 2*k {
		if len(s)-i >= k {
			ss := revers(s[i : i+k])
			s = s[:i] + ss + s[i+k:]
		} else {
			ss := revers(s[i:])
			s = s[:i] + ss
		}
	}
	return s
}

func revers(s string) string {
	bytes := []byte(s)
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
