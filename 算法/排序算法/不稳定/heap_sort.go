package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/24 6:08 下午
 */

/*
	时间复杂度：
		nlog(n)
	操作：在堆的数据结构中，堆中的最大值总是位于根节点（在优先队列中使用堆的话堆中的最小值位于根节点）。堆中定义以下几种操作：
		最大堆调整（Max Heapify）：将堆的末端子节点作调整，使得子节点永远小于父节点
		创建最大堆（Build Max Heap）：将堆中的所有数据重新排序
		堆排序（HeapSort）：移除位在第一个数据的根节点，并做最大堆调整的递归运算
	算法：
		堆排序的方法如下，把最大堆堆顶的最大数取出，
		将剩余的堆继续调整为最大堆，再次将堆顶的最大数取出，这个过程持续到剩余数只有一个时结束。
*/

func heapSort(arr []int, n int) {

	for i := n/2-1; i >= 0; i -- {
		heapify(arr,n,i)
	}
	// 一个个从堆顶取出元素
	for i := n - 1; i >= 0; i-- {
		arr[0],arr[i] = arr[i],arr[0]
		heapify(arr,i,0)
	}
}
// 最大堆调整
func heapify(arr []int, n, i int) {
	largest := i // 将最大元素设值为堆顶
	l := 2*i + 1 // 左孩子
	r := 2*i + 2 // 右孩子

	if l < n && arr[l] > arr[largest]{ // 如果左孩子比父节点大的话
		largest = l // largest "指向"左孩子
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i],arr[largest] = arr[largest],arr[i] // 交换堆顶元素和子节点
		heapify(arr,n,largest) // 递归定义子堆
	}
}

func main() {
	nums := []int{1, 4, 5, 67, 2, 7, 8, 6, 9, 44}
	heapSort(nums,len(nums))
	fmt.Println(nums)
}
