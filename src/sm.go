package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	var input string
	var sum int
	var stnum, conum float64 = 0, 2
	stack := list.New()

	fmt.Printf("请输入一段二进制数字:")
	fmt.Scanf("%s", &input)
	for _, c := range input {
		//入栈
		stack.PushBack(c)
	}

	// 出栈
	for e := stack.Back(); e != nil; e = e.Prev() {
		v := e.Value.(int32)
		sum += int(v-48) * int(math.Pow(conum, stnum))
		stnum++
	}
	fmt.Printf("二进制转化为十进制结果是 %d\n", sum)
}
