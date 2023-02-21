package 迭代

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/21
 * Time: 23:19
 * Description: No Description
 */

func subsets(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})

	for _, num := range nums {
		n := len(res)
		for i := 0; i < n; i++ {
			subset := make([]int, len(res[i]))
			copy(subset, res[i])
			subset = append(subset, num)
			res = append(res, subset)
		}
	}

	return res
}
