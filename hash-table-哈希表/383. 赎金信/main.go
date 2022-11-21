package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/21
 * Time: 23:54
 * Description: No Description
 */

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	cnt := make([]int, 26)

	for _, v := range ransomNote {
		cnt[v-'a']++
	}

	for _, v := range magazine {
		cnt[v-'a']--
		if cnt[v-'a'] < 0 {
			return false
		}
	}
	return true
}
