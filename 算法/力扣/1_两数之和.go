package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/17 10:19 上午
 */

func main() {
	nums := []int{1,2,3,4,5,6,7,8}
	target := 12
	sum := twoSum(nums, target)
	fmt.Println(sum)
}

// 暴力破解
func twoSum(nums []int, target int) []int {
	len := len(nums)
	for i := 0; i < len ; i++ {
		for j := 1; j < len; j++ {
			if (nums[i] + nums[j]) == target {
				return []int{nums[i],nums[j]}
			}
		}
	}
	return 	[]int{}
}
