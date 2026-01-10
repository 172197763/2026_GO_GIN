package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(arr []any) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	root := &TreeNode{Val: arr[0].(int)}
	i := 1
	queue := []*TreeNode{root}
	for {
		tmp_queue := []*TreeNode{}
		for _, v := range queue {
			if i >= len(arr) {
				break
			}
			if arr[i] != nil {
				v.Left = &TreeNode{Val: arr[i].(int)}
				tmp_queue = append(tmp_queue, v.Left)
			}
			i++
			if i >= len(arr) {
				break
			}
			if arr[i] != nil {
				v.Right = &TreeNode{Val: arr[i].(int)}
				tmp_queue = append(tmp_queue, v.Right)
			}
			i++
		}
		if len(tmp_queue) == 0 {
			break
		}
		queue = make([]*TreeNode, len(tmp_queue))
		copy(queue, tmp_queue)
	}
	return root
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
