package main

import (
	"fmt"
)

/**
* @Author : awen
* @Date : 2020/2/24 12:40 上午
 */

/*
	快速排序
	简称快排，时间复杂度并不固定，
		最坏 O(n^2)（和选择排序一个效率），
		比较理想 O(nlogn)。
	快排也是一个分治的算法，快排算法每次选择一个元素并且将整个数组以那个元素分为两部分，
	根据实现算法的不同，元素的选择一般有如下几种：
		+ 永远选择第一个元素
		+ 永远选择最后一个元素
		+ 随机选择元素
		+ 取中间值
	整个快速排序的核心是分区（partition），
		分区的目的是传入一个数组和选定的一个元素，把所有小于那个元素的其他元素放在左边，大于的放在右边。
	分治思想

*/
//func partition(nums []int, left,right int) int {
//	// 选取支点
//	pivot := nums[left]
//}
func quickSort(nums []int,left int ,right int) {
	// 退出递归条件
	if left >= right {
		return
	}

	l,r := left,right
	// 支点
	pivot := nums[(left+right)/2]
	for l <= r {
		// 从左往右找比 支点大的数
		for pivot > nums[l] {
			l ++
		}
		for pivot < nums[r] {
			r --
		}
		// 已经找到比支点小的数(在右边)，比支点大的数(在左边)
		if l <= r {
			// 交换左右两个数
			nums[l],nums[r] = nums[r],nums[l]
			l ++
			r --
		}
	}
	//上面保证了第一趟排序支点的左边)支点小，支点的右边比支点大了
	// 递归排序左半边
	if left < r {
		quickSort(nums,left,l)
	}
	// 递归排序右半边
	if right > l {
		quickSort(nums,l,right)
	}
}

func main() {
	nums := []int{1, 4, 5, 67, 2, 7, 8, 6, 9, 44}
	quickSort(nums,0,len(nums)-1)
	fmt.Println(nums)
}
