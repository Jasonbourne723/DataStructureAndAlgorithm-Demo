package leetcodelearn

import "fmt"

/**
 * Definition for a binary tree node.

 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var r []string
	var p string
	n := root

	var backtrack func(int)

	backtrack = func(index int) {
		if index == 0 {
			s := fmt.Sprintf("%d", n.Val)
			p += s
			index++
		}
		if n.Left == nil && n.Right == nil {
			r = append(r, p)
			return
		}

		parent := n
		if n.Left != nil {
			n = n.Left
			s := fmt.Sprintf("->%d", n.Val)
			p += s
			backtrack(index)
			p = p[:len(p)-len(s)]
			n = parent
		}

		if n.Right != nil {
			n = n.Right
			s := fmt.Sprintf("->%d", n.Val)
			p += s
			backtrack(index)
			p = p[:len(p)-len(s)]
			n = parent
		}

	}
	backtrack(0)

	return r
}
