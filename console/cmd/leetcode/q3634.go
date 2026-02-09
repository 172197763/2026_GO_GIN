package leetcode

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/minimum-removals-to-balance-array/?envType=daily-question&envId=2026-02-06
func Q3634() {
	arr := []int{20, 5, 11}
	fmt.Println(minRemoval(arr, 2))
}
func minRemoval(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	}
	res := 1 << 32
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		j := len(nums) - 1
		for i < j {
			if k*nums[i] >= nums[j] {
				break
			}
			j--
		}
		if len(nums)-j-1+i < res {
			res = len(nums) - j - 1 + i
		}
	}
	return res
}
