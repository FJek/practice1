package main

/**
* @Author : awen
* @Date : 2020/2/18 2:31 下午
 */

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	low := head
	for fast != nil {
		fast = fast.Next
		if k == 0 {
			low = low.Next
		} else {
			k --
		}
	}
	return low
}

func main() {

}
