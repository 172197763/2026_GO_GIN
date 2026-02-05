package leetcode

import "fmt"

// https://leetcode.cn/problems/trionic-array-ii/description/?envType=daily-question&envId=2026-02-04
func Q3640() {
	arr := []int{0, -2, -1, -3, 0, 2, -1}
	fmt.Println(maxSumTrionic(arr))
}
func maxSumTrionic(nums []int) int64 {
	sIdx := -1
	mid := int64(0)
	max := int64(-100000)
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 0 {
			if sIdx != -1 {
				tmpMax, end, err := cal(&nums, sIdx, i-1)
				if err == nil {
					tmpMax += mid
					if tmpMax > max {
						max = tmpMax
					}
					i = end
				}
				sIdx = -1
				mid = 0
			}
		} else if nums[i]-nums[i-1] < 0 {
			if sIdx == -1 {
				sIdx = i
			}
			mid += int64(nums[i])
		}
	}
	return max
}
func cal(nums *[]int, start, end int) (int64, int, error) {
	if start == 0 || end == len(*nums)-1 {
		return 0, 0, nil
	}
	lMax := -100000
	rMax := -100000
	tmpMax := 0
	for {
		start--
		tmpMax += (*nums)[start]
		if lMax < tmpMax {
			lMax = tmpMax
		}
		if start == 0 || (*nums)[start-1] > (*nums)[start] {
			break
		}
	}

	tmpMax = 0
	for {
		end++
		tmpMax += (*nums)[end]
		if rMax < tmpMax {
			rMax = tmpMax
		}

		if end == len(*nums)-1 || (*nums)[end] > (*nums)[end+1] {
			break
		}
	}

	return int64(lMax + rMax), end, nil
}
