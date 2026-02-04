package leetcode

import "fmt"

// https://leetcode.cn/problems/trionic-array-i/?envType=daily-question&envId=2026-02-03
func Q3637() {
	arr := []int{1, 3, 5, 4, 2, 6}
	fmt.Println(isTrionic(arr))
}
func isTrionic(nums []int) bool {
	clean := 0
	bit := 0
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 0 {
			if clean != 1 {
				clean = 1
				res = res | 1<<bit
				bit++
			}
		} else if nums[i]-nums[i-1] < 0 {
			if clean != -1 {
				clean = -1
				bit++
			}
		} else {
			return false
		}
	}
	return res == 5 && bit == 3
}
