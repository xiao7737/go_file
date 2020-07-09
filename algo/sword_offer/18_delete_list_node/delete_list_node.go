package _8_delete_list_node

//删除链表的节点，返回处理后的链表头结点
type ListNode struct {
	Val  int
	Next *ListNode
}

//双指针迭代，时间O(N)，空间O(1)
func deleteNode(head *ListNode, val int) *ListNode {
	//删除节点为head，则直接返回head.Next
	if head.Val == val {
		return head.Next
	}
	var pre *ListNode
	//初始化pre和head
	pre = head
	cur := head.Next

	for {
		//修改引用，删除节点
		if cur.Val == val {
			pre.Next = cur.Next
			break
		}
		//向后遍历
		pre = cur
		cur = cur.Next
	}
	return head
}

//递归 时间O(N)，空间O(N)
func deleteNode1(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	if head.Next != nil {
		head.Next = deleteNode1(head.Next, val)
	}
	return head
}
