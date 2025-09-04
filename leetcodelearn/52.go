package leetcodelearn

//n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。

func totalNQueens(n int) int {
	var r int

	x := make([]int, n)
	y := make([]int, n)
	used := [][2]int{}

	var backtrack func(int)

	backtrack = func(row int) {
		if len(used) == n {
			r++
			return
		}
	loop:
		for col := 0; col < n; col++ {
			if row < 0 || row >= n {
				break
			}
			if x[row] == 1 {
				continue
			}
			if y[col] == 1 {
				continue
			}

			for _, item := range used {
				if item[0]-row == item[1]-col || item[0]-row == col-item[1] {
					continue loop
				}
			}
			x[row] = 1
			y[col] = 1
			used = append(used, [2]int{row, col})

			backtrack(row + 1)

			used = used[:len(used)-1]
			x[row] = 0
			y[col] = 0
		}
	}

	backtrack(0)
	return r
}
