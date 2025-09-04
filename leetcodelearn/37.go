package leetcodelearn

//编写一个程序，通过填充空格来解决数独问题。
//
//数独的解法需 遵循如下规则：
///数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//数独部分空格内已填入了数字，空白格用 '.' 表示。

func SolveSudoku(board [][]byte) {
	x := [9][10]bool{}
	y := [9][10]bool{}
	z := [9][10]bool{}
	var empty_count int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				x[i][board[i][j]-48] = true
				y[j][board[i][j]-48] = true
				z[i/3*3+j/3][board[i][j]-48] = true
			} else {
				empty_count++
			}
		}
	}

	var backtrack func(int) bool

	backtrack = func(index1 int) bool {

		if empty_count == 0 {
			return true
		}

		if index1 < 0 || index1 >= 81 {
			return false
		}
		i := index1 / 9
		j := index1 % 9

		if board[i][j] != '.' {
			return backtrack(index1 + 1)
		}

		for num := 1; num <= 9; num++ {

			if x[i][num] {
				continue
			}
			if y[j][num] {
				continue
			}
			// 小方形
			if z[i/3*3+j/3][num] {
				continue
			}

			board[i][j] = byte(num + 48)
			x[i][num] = true
			y[j][num] = true
			z[i/3*3+j/3][num] = true
			empty_count--

			if backtrack(index1 + 1) {
				return true
			}

			board[i][j] = '.'
			x[i][num] = false
			y[j][num] = false
			z[i/3*3+j/3][num] = false
			empty_count++

		}
		return false
	}
	backtrack(0)
}
