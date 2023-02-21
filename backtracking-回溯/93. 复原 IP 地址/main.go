package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/21
 * Time: 21:33
 * Description: https://leetcode.cn/problems/restore-ip-addresses/
 */

func main() {
	s := "25525511135"
	ips := restoreIpAddresses(s)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)

	backtrack(s, &res, []string{}, 0)
	return res
}

func backtrack(s string, res *[]string, path []string, start int) {
	if len(path) == 4 {
		if start == len(s) { // 这里说明这一层已经遍历完了
			str := strings.Join(path, ".")
			*res = append(*res, str)
		}
		return
	}

	for i := start; i < len(s); i++ {
		if i != start && s[start] == '0' { // 含有前导0， 无效
			break
		}
		str := s[start : i+1]
		num, _ := strconv.Atoi(str)
		if num >= 0 && num <= 255 {
			path = append(path, str) // 符合条件的就进入下一层
			backtrack(s, res, path, i+1)
			path = path[:len(path)-1]
		} else { // 如果不满足条件，再往后也不可能满足条件，直接退出
			break
		}
	}
}
