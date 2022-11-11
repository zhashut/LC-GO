package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/11
 * Time: 23:09
 * Description: No Description
 */

func main() {
	fruits := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	totalFruit(fruits)
}

// 滑动窗口
func totalFruit(fruits []int) (ans int) {
	cnt := map[int]int{}
	left := 0

	for right, val := range fruits {
		cnt[val]++
		for len(cnt) > 2 {
			y := fruits[left]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
