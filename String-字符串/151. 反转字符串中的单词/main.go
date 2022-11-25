package main

import "strings"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/26
 * Time: 0:16
 * Description: No Description
 */

func main() {
	s := "  hello world  "
	reverseWords(s)
}

func reverseWords(s string) string {
	// 1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	// 删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	// 删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	// 删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	// 2.反转整个字符串
	reverse(&b, 0, len(b)-1)
	// 3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}

// 解法，方法解题
func reverseWords1(s string) string {
	ss := strings.Fields(s)
	reverse1(&ss, 0, len(ss)-1)
	return strings.Join(ss, " ")
}

func reverse1(m *[]string, i int, j int) {
	for i <= j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		i++
		j--
	}
}
