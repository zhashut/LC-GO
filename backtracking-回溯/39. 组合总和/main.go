package _9__组合总和

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/19
 * Time: 22:25
 * Description: https://leetcode.cn/problems/combination-sum/
 */

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	backtrack(&result, []int{}, candidates, target, 0)
	return result
}

func backtrack(result *[][]int, tempList []int, candidates []int, remain int, start int) {
	if remain < 0 {
		return
	}
	if remain == 0 {
		// make a copy of the current list and append it to the result
		temp := make([]int, len(tempList))
		copy(temp, tempList)
		*result = append(*result, temp)
	}
	for i := start; i < len(candidates); i++ {
		tempList = append(tempList, candidates[i])
		backtrack(result, tempList, candidates, remain-candidates[i], i)
		tempList = tempList[:len(tempList)-1]
	}
}
