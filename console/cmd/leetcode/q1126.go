package leetcode

import (
	"fmt"
)

// https://leetcode.cn/problems/smallest-subtree-with-all-the-deepest-nodes/?envType=daily-question&envId=2026-01-09
func Q1126() {
	arr := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	fmt.Println(minTimeToVisitAllPoints(arr))
}
func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	for i := 0; i < len(points)-1; i++ {
		res += max(abs(points[i][0]-points[i+1][0]), abs(points[i][1]-points[i+1][1]))
	}
	return res
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
