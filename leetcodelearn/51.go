package leetcodelearn

/*按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。*/

func SolveNQueens(n int) [][]string {
	var r [][]string
	path := make([][]byte, n)
	x := make([]byte, n)
	y := make([]byte, n)
	ex := make([][2]int, n)
	var q int
	for i := 0; i < n; i++ {
		bs := make([]byte, n)
		for j := 0; j < n; j++ {
			bs[j] = '.'
		}
		path[i] = bs
	}
	var backtrack func(int)

	backtrack = func(i int) {
		if q == n {
			t := make([]string, n)
			for index, item := range path {
				t[index] = string(item)
			}
			r = append(r, t)
			return
		}

	loop:
		for j := 0; j < n; j++ {
			if i >= n {
				break
			}
			if x[i] == '1' {
				break
			}
			if y[j] == '1' {
				continue
			}

			if q != 0 {
				for k := 0; k < q; k++ {
					a := ex[k][0] - i
					b := ex[k][1] - j

					if a == b || a+b == 0 {
						continue loop
					}
				}
			}

			x[i] = '1'
			y[j] = '1'
			path[i][j] = 'Q'
			ex[q] = [2]int{i, j}
			q++
			backtrack(i + 1)
			path[i][j] = '.'
			x[i] = '0'
			y[j] = '0'
			q--
		}
	}

	for i := 0; i < n; i++ {
		backtrack(i)
	}
	return r
}
