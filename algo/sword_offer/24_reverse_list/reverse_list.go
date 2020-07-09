package _4_reverse_list

//翻转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

//双指针迭代   时间O(N), 空间O(1)
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		//往后迭代
		pre = cur
		cur = temp
	}
	return pre
}

//递归 	时间O(N), 空间O(N)
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList1(head.Next)
	head.Next.Next = head
	//终止本次递归，防止链表循环
	head.Next = nil
	return p
}
