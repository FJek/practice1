package main

/**
* @Author : awen
* @Date : 2020/2/18 2:41 下午
 */

// 反转链表
// 快慢指针
func reverseList(head *ListNode) *ListNode {
	// fast 指针在low指针前面
	var low *ListNode
	pre := head
	for pre != nil {
		temp := pre.Next // 临时存放 pre 的下一个，避免断开链表
		pre.Next = low
		low = pre
		pre = temp
	}
	return low
}

// 递归
func reverseList2(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(pre,cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	head := reverse(cur, cur.Next)
	cur.Next = pre
	return head
}
