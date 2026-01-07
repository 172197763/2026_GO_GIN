package leetcode

import (
	"fmt"
	"slices"
)

// https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/?envType=daily-question&envId=2026-01-06
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Q1161() {
	root := &TreeNode{989, nil, nil}
	root.Right = &TreeNode{10250, nil, nil}
	root.Right.Left = &TreeNode{98693, nil, nil}
	root.Right.Right = &TreeNode{-89388, nil, nil}
	root.Right.Right.Right = &TreeNode{-32127, nil, nil}
	fmt.Println(maxLevelSum(root))
}

// result:✅️pass
func maxLevelSum(root *TreeNode) int {
	maxNums := root.Val
	maxF := 1
	queue := []*TreeNode{root}
	f := 2
	for len(queue) > 0 {
		tmp_queue := []*TreeNode{}
		tmp_max := 0
		for _, v := range queue {
			if v.Left != nil {
				tmp_max += v.Left.Val
				tmp_queue = append(tmp_queue, v.Left)
			}
			if v.Right != nil {
				tmp_max += v.Right.Val
				tmp_queue = append(tmp_queue, v.Right)
			}
		}
		if len(tmp_queue) == 0 {
			break
		}
		if tmp_max > maxNums {
			maxNums = tmp_max
			maxF = f
		}
		f++
		queue = slices.Clone(tmp_queue)
	}

	return maxF
}
