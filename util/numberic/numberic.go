package numberic

/**
* @Author : awen
* @Date : 2020/2/18 2:00 下午
 */

// 交换数组的两个数
func Swap(nums []int,index1 int, index2 int) {
	temp := nums[index1]
	nums[index1] = nums[index2]
	nums[index2] = temp
}
