package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/24
 * Time: 22:51
 * Description: No Description
 */

func main() {

}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseString1(s []byte) {
	rever := []byte{}

	for _, val := range s {
		rever = append(rever, val)
	}

	for i := 0; i < len(s); i++ {
		s[i] = rever[len(rever)-1-i]
	}

	return
}
