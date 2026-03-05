package leetcode

import (
	"fmt"
)

// https://leetcode.cn/problems/special-positions-in-a-binary-matrix/description/?envType=daily-question&envId=2026-03-04
func Q1758() {
	fmt.Println(minOperations("0100"))
}
func minOperations(s string) int {
	res1 := 0
	res2 := 0
	for k, v := range s {
		// fmt.Printf("%T:%d \n", v, v-48)
		// fmt.Printf("%d \n", k%2)
		if k%2 == 0 {
			if v-48 == 1 {
				res1++
			} else {
				res2++
			}
		} else {
			if v-48 == 1 {
				res2++
			} else {
				res1++
			}
		}
	}
	if res1 > res2 {
		return res2
	} else {
		return res1
	}
}
