package main

import (
	"container/heap"
	"fmt"
)

/**
* 寻找无序数组的中位数
* @Author : awen
* @Date : 2020/1/19 8:03 下午
 */


// 寻找一个无序数组的中位数
// https://zhuanlan.zhihu.com/p/83315438
func findMedianNum(objs []Obj) int {
	hp := &RectHeap{}
	heap.Init(hp)
	n := len(objs)
	k := n/2+1
	// 建堆
	for i := 0; i < k; i++ {
		hp.Push(objs[i].Num)
	}
	for i := k; i < n; i++ {
		if heap.Pop(hp).(int) < objs[i].Num {
			heap.Remove(hp,i)
			heap.Push(hp,objs[i].Num)
		}
	}
	if n%2 == 0 {
		return (heap.Pop(hp).(int)+heap.Pop(hp).(int)) /2
	} else {
		return heap.Pop(hp).(int)
	}
}



//定义一个正方形的结构体
type Obj struct {
	Num int
}



// 定义一个堆结构体
type RectHeap []Obj

// 实现heap.Interface接口
func (rech RectHeap) Len() int {
	return len(rech)
}

// 实现sort.Iterface
func (rech RectHeap) Swap(i, j int) {
	rech[i], rech[j] = rech[j], rech[i]
}
func (rech RectHeap) Less(i, j int) bool {
	return rech[i].Num < rech[j].Num
}

// 实现heap.Interface接口定义的额外方法
func (rech *RectHeap) Push(h interface{}) {
	*rech = append(*rech, h.(Obj))
}
func (rech *RectHeap) Pop() (x interface{}) {
	n := len(*rech)
	x = (*rech)[n-1]      // 返回删除的元素
	*rech = (*rech)[:n-1] // [n:m]不包括下标为m的元素
	return x
}

func main() {
	hp := &RectHeap{}
	for i := 2; i < 6; i++ {
		*hp = append(*hp, Obj{ i})
	}

	fmt.Println("原始slice: ", hp)

	// 堆操作
	heap.Init(hp)
	heap.Push(hp, Obj{10})
	fmt.Println("top元素：", (*hp)[0])
	fmt.Println("删除并返回最后一个：", heap.Pop(hp)) // 最后 一个元素
	fmt.Println("最终slice: ", hp)
}
