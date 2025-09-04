package leetcodelearn

func Exist(board [][]byte, word string) bool {
	row := len(board)
	column := len(board[0])

	v := make([][]bool, row)
	for i := range v {
		v[i] = make([]bool, column)
	}
	s := make([][]int, row)
	for i := range v {
		s[i] = make([]int, column)
	}
	var s1 int
	var path []byte

	var backtrack func(i, j int) bool
	backtrack = func(i, j int) bool {
		if i < 0 || i >= row || j < 0 || j >= column {
			return false
		}
		if v[i][j] {
			return false
		}
		v[i][j] = true
		s1++
		s[i][j] = s1
		path = append(path, board[i][j])

		if string(path) == word {
			println(string(path))
			return true
		}
		if !match(word, path) {
			v[i][j] = false
			s1--
			s[i][j] = 0
			path = path[:len(path)-1]
			return false
		}

		if backtrack(i+1, j) || backtrack(i-1, j) || backtrack(i, j+1) || backtrack(i, j-1) {
			return true
		}
		v[i][j] = false
		path = path[:len(path)-1]
		s1--
		s[i][j] = 0
		return false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if backtrack(i, j) {
				return true
			}
		}
	}
	return false
	//re := backtrack(0, 0)
	//green := color.New(color.FgGreen, color.Bold)
	//red := color.New(color.FgRed, color.Bold)
	//for i, item1 := range board {
	//	for j, item2 := range item1 {
	//		if v[i][j] {
	//			green.Print(string(item2))
	//		} else {
	//			red.Print(string(item2))
	//		}
	//		print(" ")
	//	}
	//	print("\n")
	//}
	//
	//for i, item1 := range s {
	//	for j, item2 := range item1 {
	//		if v[i][j] {
	//			green.Print(item2)
	//		} else {
	//			red.Print(item2)
	//		}
	//		print(" ")
	//	}
	//	print("\n")
	//}
	//return re
}

func match(word string, s []byte) bool {
	if len(s) > len(word) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != word[i] {
			return false
		}
	}
	return true
}
