package main

/**
* @Author : awen
* @Date : 2020/2/17 10:28 上午
 */

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	carry := 0; sum := 0 //
	p:=l1;q:=l2
	cur := head
	for p != nil || q != nil {  // 222+ 394
		x,y := 0,0 // 如果长度不等， 则 Next.Val = 0
		if p != nil {
			x = p.Val  // x = 2
			p = p.Next // 指向下一位
		}
		if q != nil {
			y = q.Val // y = 3_工厂方法
			q = q.Next
		}
		sum = x + y + carry  // sum = 5
		cur.Next = &ListNode{Val:  sum%10}
		cur = cur.Next
		carry = sum/10
	}
	if carry > 0 {
		cur.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}
