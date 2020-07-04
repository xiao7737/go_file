package _6_reserve_print

//从尾到头，打印链表，输出数组
type ListNode struct {
	Val  int
	Next *ListNode
}

//利用栈
func reservePrint1(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	stack := []int{}
	//从head开始遍历，将数据装入stack中
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	//将stack中的数据首尾交换
	/*i, j := 0, len(stack)-1
	for i < j {
		stack[i], stack[j] = stack[j], stack[i]
		i++
		j--
	}*/
	//少申请一个变量
	for i := 0; i < len(stack)>>1; i++ {
		stack[i], stack[len(stack)-1-i] = stack[len(stack)-1-i], stack[i]
	}
	return stack
}

//利用递归
//时间空间都是O(N)
func reservePrint2(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	res := reservePrint2(head.Next)
	return append(res, head.Val)
}
