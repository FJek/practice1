package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/24 5:37 下午
 */

/* 选择排序
	思想：
		首先在未排序的数列中找到最小(or最大)元素，然后将其存放到数列的起始位置；
		接着，再从剩余未排序的元素中继续寻找最小(or最大)元素，然后放到已排序序列的末尾。
		以此类推，直到所有元素均排序完毕。
	复杂度：
		n^2
*/

func selectSort(nums []int,n int)  {
	var min int // 无序区中最小元素位置
	for i := range nums {
		for j := i+1; j < n; j++ {
			if nums[j] < nums[i] {
				min = j
			}
		}
		if min != i {
			nums[i],nums[min] = nums[min],nums[i]
		}
	}
}

func main() {
	nums := []int{1, 4, 5, 67, 2, 7, 8, 6, 9, 44}
	selectSort(nums,len(nums))
	fmt.Println(nums)
}
