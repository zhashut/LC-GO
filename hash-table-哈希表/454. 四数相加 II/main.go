package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/20
 * Time: 21:02
 * Description: No Description
 */

func main() {
	a, b, c, d := []int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}
	fourSumCount(a, b, c, d)
}

func fourSumCount(a []int, b []int, c []int, d []int) (ans int) {
	cnt := map[int]int{}
	for _, i := range a {
		for _, j := range b {
			cnt[i+j]++
		}
	}

	for _, i := range c {
		for _, j := range d {
			ans += cnt[-i-j]
		}
	}

	return
}
