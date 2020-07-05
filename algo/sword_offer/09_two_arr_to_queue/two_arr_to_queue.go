package _9_two_arr_to_queue

//使用两个栈实现队列
//时间O(1)  每个元素最多移动两次
//空间O(N)  两个栈空间
//可以利用list容器，官解使用的list
type CQueue struct {
	stackA []int
	stackB []int
}

func Constructor() CQueue {
	return CQueue{
		stackA: make([]int, 0),
		stackB: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stackA = append(this.stackA, value)
}

func (this *CQueue) DeleteHead() int {
	//判断stackB是否为空
	if len(this.stackB) == 0 {
		//将stackA数据移动到stackB ，逐渐将stackA的最后一个元素装入stackB
		for len(this.stackA) > 0 {
			this.stackB = append(this.stackB, this.stackA[len(this.stackA)-1])
			this.stackA = this.stackA[:len(this.stackA)-1]
		}
	}
	if len(this.stackB) != 0 {
		//stackB最后一个元素就是要删除的元素
		deleteEle := this.stackB[len(this.stackB)-1]
		this.stackB = this.stackB[:len(this.stackB)-1]
		return deleteEle
	}
	return -1
}
