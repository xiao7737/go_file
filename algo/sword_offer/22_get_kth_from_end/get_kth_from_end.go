package _2_get_kth_from_end

//获取链表倒数第k个节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间O(N)，first节点走了N步
// 空间O(1)
func getKthFromEnd(head *ListNode, k int) *ListNode {
	first, last := head, head
	//fist节点先走k步
	for i := 0; i < k; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		last = last.Next
	}
	//直到first走到链表尾部，返回last，此时last是从last开始的链表
	return last
}
