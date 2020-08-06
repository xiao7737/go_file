package _2_get_first_node_in_2_list

//返回两个链表的第一个公共节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//利用map，存地址和值，如果相同，则是重复的第一个公共节点
//时间O(N)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]int)
	for headA != nil || headB != nil {
		if headA != nil {
			if _, ok := m[headA]; !ok {
				m[headA] = headA.Val
			} else {
				//m中已经有这个值了，则表示该节点重复了，返回即可
				return headA
			}
			//指向下一个
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := m[headB]; !ok {
				m[headB] = headB.Val
			} else {
				return headB
			}
			//指向下一个
			headB = headB.Next
		}
	}
	return nil
}

//todo 可以参考答案写法
