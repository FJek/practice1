package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/28 10:38 下午
 */


func main() {
	arr := []int{1 ,1 ,1 ,2 ,2 ,2, 3, 3 ,3, 3}
	res := noduplicate(arr)
	fmt.Println(res)
}

func noduplicate(arr []int) int {
	last := arr[0]
	count := 1 //当前数字的个数
	max := 0 // 重复最多的次数
	for i := 1;i<len(arr);i++ {
		//fmt.Println(i,v)
		if arr[i] == last {
			count += 1
			last = arr[i]
		} else {
			if count > max{
				max = count
			}
			count = 1
			last = arr[i]
		}
	}
	// 最后一个重复序列也要比较
	if count > max{
		max = count
	}
	return len(arr)-max
}
