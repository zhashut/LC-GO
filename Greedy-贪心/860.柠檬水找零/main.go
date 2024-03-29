package _60_柠檬水找零

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/4
 * Time: 22:29
 * Description: https://leetcode.cn/problems/lemonade-change/
 */

func lemonadeChange(bills []int) bool {
	five, ten, twenty := 0, 0, 0

	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
		}
		if bills[i] == 10 {
			if five <= 0 {
				return false
			}
			five--
			ten++
		}

		if bills[i] == 20 {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
				twenty++
			} else {
				return false
			}
		}
	}

	return true
}
