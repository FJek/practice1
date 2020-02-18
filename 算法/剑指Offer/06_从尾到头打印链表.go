package main

import (
	"container/list"
	"fmt"
)

/**
* @Author : awen
* @Date : 2020/2/18 2:06 下午
 */


type ListNode struct {
	Val int
	Next *ListNode
}

// 从尾到头打印链表
func ReversePrint(head *ListNode) []int {
	res := make([]int, 0)
	l := list.New() // 创建一个链表
	for ; head != nil; head = head.Next {
		// 遍历每一个节点
		l.PushFront(head.Val) // 插入到新链表的头部
	}

	for item := l.Front(); item != nil; item = item.Next(){
		res = append(res, item.Value.(int))
		fmt.Println(item.Value)
	}
	return res
}

func main() {

}
