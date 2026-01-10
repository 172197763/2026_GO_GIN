package leetcode

import (
	"fmt"
	"slices"
)

// https://leetcode.cn/problems/smallest-subtree-with-all-the-deepest-nodes/?envType=daily-question&envId=2026-01-09
func Q865() {
	arr := []any{0, 1, nil, 3, 2, 6, nil, 5, 4}
	root := buildTree(arr)
	PrintTreeStructure(root)
	fmt.Println(*subtreeWithAllDeepest(root))
}

var total = 0

// result:✅️pass
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	f := []*TreeNode{root}
	p_map := map[*TreeNode]*TreeNode{}
	for {
		tmp_f := []*TreeNode{}
		for _, v := range f {
			if v.Left != nil {
				p_map[v.Left] = v
				tmp_f = append(tmp_f, v.Left)
			}
			if v.Right != nil {
				p_map[v.Right] = v
				tmp_f = append(tmp_f, v.Right)
			}
		}
		if len(tmp_f) == 0 {
			if len(f) == 1 {
				return f[0]
			}
			break
		}
		f = make([]*TreeNode, len(tmp_f))
		copy(f, tmp_f)
	}
	for {
		tmp_f := []*TreeNode{}
		for _, v := range f {
			if !slices.Contains(tmp_f, p_map[v]) {
				tmp_f = append(tmp_f, p_map[v])
			}
		}
		f = make([]*TreeNode, len(tmp_f))
		copy(f, tmp_f)
		if len(tmp_f) == 1 {
			break
		}
	}
	return f[0]
}
