package reverseList_206

// 反转链表 1 2 3 4 5   >>>>   5 4 3 2 1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针迭代
// 时间O(N), 空间O(1)
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

// 递归
// 时间O(N), 空间O(N)
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//p为最后一个需要交换的节点
	p := reverseList1(head.Next)

	// 3.next.next=3   相当于4.next=3
	head.Next.Next = head

	//防止链表循环
	head.Next = nil

	return p
}
