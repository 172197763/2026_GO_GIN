package leetcode

import (
	"fmt"
)

// https://leetcode.cn/problems/maximum-product-of-splitted-binary-tree/description/?envType=daily-question&envId=2026-01-07

func Q1339() {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{4, nil, nil}
	root.Left.Right = &TreeNode{5, nil, nil}
	maxProduct(root)
}

// result:✅️pass
func maxProduct(root *TreeNode) int {
	dfs1339(root)
	res := dfscal1339(root, root.Val)
	mod := int64(1000000000 + 7)
	return int(res % mod)
}
func dfscal1339(node *TreeNode, max int) int64 {
	res := int64((max - node.Val) * node.Val)
	if node.Left != nil {
		left := dfscal1339(node.Left, max)
		if left > res {
			res = left
		}
	}
	if node.Right != nil {
		right := dfscal1339(node.Right, max)
		if right > res {
			res = right
		}
	}
	return res
}
func dfs1339(node *TreeNode) int {
	sum := node.Val
	if node.Left != nil {
		sum += dfs1339(node.Left)
	}
	if node.Right != nil {
		sum += dfs1339(node.Right)
	}
	node.Val = sum
	return node.Val
}
func PrintTreeStructure(root *TreeNode) {
	if root == nil {
		fmt.Println("Empty tree")
		return
	}
	printTreeHelper(root, "", true)
}

func printTreeHelper(node *TreeNode, prefix string, isLast bool) {
	if node != nil {
		fmt.Print(prefix)
		if isLast {
			fmt.Print("└── ")
			prefix += "    "
		} else {
			fmt.Print("├── ")
			prefix += "│   "
		}
		fmt.Println(node.Val)

		// 计算子节点数量
		children := 0
		if node.Left != nil {
			children++
		}
		if node.Right != nil {
			children++
		}

		// 递归打印左子树
		if node.Left != nil {
			printTreeHelper(node.Left, prefix, children == 1 && node.Right == nil)
		}

		// 递归打印右子树
		if node.Right != nil {
			printTreeHelper(node.Right, prefix, true)
		}
	}
}
