package main

import "fmt"

/**
 * Created with GoLand 2022.2.3.
 * @author: zhashut
 * Date: 2025/2/26
 * Time: 20:11
 * Description: https://leetcode.cn/problems/design-browser-history/description/
 */

func main() {
	history := Constructor("leetcode.com")
	history.Visit("google.com")
	history.Visit("facebook.com")
	history.Visit("youtube.com")
	fmt.Println(history.Back(1))
	fmt.Println(history.Back(1))
	fmt.Println(history.Forward(1))
	history.Visit("linked.com")
	fmt.Println(history.Forward(2))
	fmt.Println(history.Back(2))
	fmt.Println(history.Back(7))
}
