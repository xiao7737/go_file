package _2_get_kth_from_end

//获取链表倒数第k个节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间O(N)，first节点走了N步
// 空间O(1)
func getKthFromEnd(head *ListNode, k int) *ListNode {
	first, second := head, head
	//fist节点先走k步
	for i := 0; i < k; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	//直到first走到链表尾部，返回last
	return second
}

//从鲁棒性出发，以上代码存在问题：
//1: 输入的头指针如果为空，或者k为0，程序崩溃
//2: 如果链表的总长度小于k，会程序崩溃
func getKthFromEnd2(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return nil
	}
	first, second := head, head
	for i := 0; i < k; i++ {
		//处理链表长度小于k的情况
		if first.Next != nil {
			first = first.Next
		} else {
			return nil
		}
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	return second
}
