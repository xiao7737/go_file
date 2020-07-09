package _5_merge_two_list

//合并两个有序链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//递归做法  时间O(M+N) ，空间O(M+N)
func mergeTwoList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l2.Next = mergeTwoList(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoList(l1.Next, l2)
		return l1
	}
}
