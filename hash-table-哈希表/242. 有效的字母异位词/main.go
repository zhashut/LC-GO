package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/18
 * Time: 11:39
 * Description: https://leetcode.cn/problems/valid-anagram/
 */

func main() {
	s, v := "anagram", "nagaraam"
	isAnagram(s, v)
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	cnt := map[rune]int{}

	for _, val := range s {
		cnt[val]++
	}

	for _, val := range t {
		cnt[val]--
		if cnt[val] < 0 {
			return false
		}
	}

	return true
}
