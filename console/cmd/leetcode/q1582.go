package leetcode

import (
	"fmt"
)

// https://leetcode.cn/problems/special-positions-in-a-binary-matrix/description/?envType=daily-question&envId=2026-03-04
func Q1582() {
	arr := [][]int{{1, 0, 0}, {0, 1, 0}, {1, 0, 0}}
	fmt.Println(numSpecial(arr))
}
func numSpecial(mat [][]int) int {
	col := make([]int, len(mat[0]))
	//初始化
	for k := range col {
		col[k] = -1
	}
	res := 0
	for _, v := range mat {
		oneNums := 0
		_res := 0
		for j, vv := range v {
			if vv == 1 {
				oneNums++
				if oneNums > 1 {
					break
				}
				//未检测的列则检测一次
				if col[j] == -1 {
					col[j] = 0
					for i := 0; i < len(mat); i++ {
						if mat[i][j] == 1 {
							col[j]++
						}
					}
				}
				if col[j] == 1 {
					_res++
				}
			}
		}
		if oneNums == 1 {
			res += _res
		}
	}
	return res
}
