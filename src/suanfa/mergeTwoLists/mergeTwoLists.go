package mergeTwoLists

//合并两个有序链表为一个有序链表
//结题思路：利用递归，寻找节点中比较小的节点，然后从这个这个节点开始向后走

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	//如果l2的节点比l1小，那么从l2的节点开始向后走，这样形成一个小数在前，依次递增的链表
	if l1.Val > l2.Val {
		l2.Next = mergeTwoList(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoList(l1.Next, l2)
		return l1
	}
}
