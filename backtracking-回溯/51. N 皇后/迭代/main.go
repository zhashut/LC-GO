package main

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/24
 * Time: 21:18
 * Description: https://leetcode.cn/problems/n-queens/
 */

/* 解法二 二进制操作法 https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0051.N-Queens/
还有一个位运算的方法，每行只能选一个位置放皇后，那么对每行遍历可能放皇后的位置。
如何高效判断哪些点不能放皇后呢？这里的做法毕竟巧妙，把所有之前选过的点按照顺序存下来，
然后根据之前选的点到当前行的距离，就可以快速判断是不是会有冲突。举个例子: 假如在 4 皇后问题中，
如果第一二行已经选择了位置 [1, 3]，那么在第三行选择时，首先不能再选 1, 3 列了，而对于第三行，
1 距离长度为2，所以它会影响到 -1, 3 两个列。同理，3 在第二行，距离第三行为 1，所以 3 会影响到列 2, 4。
由上面的结果，我们知道 -1, 4 超出边界了不用去管，别的不能选的点是 1, 2, 3，所以第三行就只能选 0。
在代码实现中，可以在每次遍历前根据之前选择的情况生成一个 occupied 用来记录当前这一行，
已经被选了的和由于之前皇后攻击范围所以不能选的位置，然后只选择合法的位置进入到下一层递归。
另外就是预处理了一个皇后放不同位置的字符串，
这样这些字符串在返回结果的时候是可以在内存中复用的，省一点内存。
*/

func main() {
	solveNQueens(4)
}

func solveNQueens(n int) (res [][]string) {
	placements := make([]string, n)
	for i := range placements {
		buf := make([]byte, n)
		for j := range placements {
			if i == j {
				buf[j] = 'Q'
			} else {
				buf[j] = '.'
			}
		}
		placements[i] = string(buf)
	}
	var construct func(prev []int)
	construct = func(prev []int) {
		if len(prev) == n {
			plan := make([]string, n)
			for i := 0; i < n; i++ {
				plan[i] = placements[prev[i]]
			}
			res = append(res, plan)
			return
		}
		occupied := 0
		for i := range prev {
			dist := len(prev) - i
			bit := 1 << prev[i]
			occupied |= bit | bit<<dist | bit>>dist
		}
		prev = append(prev, -1)
		for i := 0; i < n; i++ {
			if (occupied>>i)&1 != 0 {
				continue
			}
			prev[len(prev)-1] = i
			construct(prev)
		}
	}
	construct(make([]int, 0, n))
	return
}
