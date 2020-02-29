package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
* @Author : awen
* @Date : 2020/2/29 11:19 上午
 */

func main() {
	// 数组多大？
	// [13,2,6,1]
	var numStr string
	fmt.Scan(&numStr)
	//// 去掉头尾【 】
	ss := numStr[1 : len(numStr)-1]
	////fmt.Println(ss)
	strArr := strings.Split(ss, ",")
	//// 转int arr
	//nums := make([]int,0)

	nums := make([]int,0)
	for _,v := range strArr{
		vInt,_ := strconv.Atoi(v)
		nums = append(nums,vInt)
	}
	//fmt.Println(nums)
	//fmt.Println(nums)
	res := "["
	counts := lessNumber(nums)
	for i,v := range counts{
		s := strconv.Itoa(v)
		if i == len(counts)-1 {
			res +=s +"]"
			break
		}
		res += s+","
	}
	fmt.Println(res)
}

func lessNumber(nums []int) []int {
	counts := make([]int,0)
	for i := 0; i < len(nums); i++ {
		size := 0
		for j := i+1; j<len(nums) ;j++  {
			if  nums[j] < nums[i]{
				size ++
			}
		}
		counts = append(counts, size)
	}
	return counts
}
