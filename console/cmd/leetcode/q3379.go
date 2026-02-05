package leetcode

import "fmt"

// https://leetcode.cn/problems/transformed-array/description/?envType=daily-question&envId=2026-02-05
func Q3379() {
	arr := []int{3, -2, 1, 1}
	//           0  1  2  3
	arr = []int{-1, 4, -1}
	fmt.Println(constructTransformedArray(arr))
}
func constructTransformedArray(nums []int) []int {
	l := len(nums) //3
	res := make([]int, l)
	for i, v := range nums {
		target := i + v%l
		if v > 0 {
			res[i] = nums[target%l]
		} else if v < 0 {
			target += l
			res[i] = nums[target%l]
		} else {
			res[i] = v
		}
	}
	return res
}
