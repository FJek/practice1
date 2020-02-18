package main

import (
	"fmt"
	"fzw/practice1/util/numberic"
)

/**
 * @Description:
 * @Date: 2020/2/18
 * @Time: 1:50 下午
 */

func findRepeatNumber(nums []int) int {
	l := len(nums)
	for i := 0; i < l; i++ { // 遍历每一个数
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			numberic.Swap(nums,i,nums[i])
		}
	}
	return 0
}



func main() {
	number := findRepeatNumber([]int{1, 2, 3,3, 4, 5, 4, 4,})
	fmt.Println(number)
}
