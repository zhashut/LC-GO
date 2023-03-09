package _38__单调递增的数字

import "strconv"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/9
 * Time: 20:23
 * Description: https://leetcode.cn/problems/monotone-increasing-digits/
 */

func monotoneIncreasingDigits(n int) int {
	nStr := strconv.Itoa(n) // 将数字转为字符串，方便使用下标
	ss := []byte(nStr)      // 将字符串转为byte数组，方便更改。
	length := len(ss)

	for i := length - 1; i > 0; i-- {
		if ss[i-1] > ss[i] { //前一个大于后一位,前一位减1，后面的全部置为9
			ss[i-1] -= 1
			for j := i; j < length; j++ { //后面的全部置为9
				ss[j] = '9'
			}
		}
	}

	res, _ := strconv.Atoi(string(ss))
	return res
}
