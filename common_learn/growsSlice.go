package main

import "fmt"

func main() {
	s := make([]int, 0)

	oldCap := cap(s)

	for i := 0; i < 2048; i++ {
		s = append(s, i)

		newCap := cap(s)

		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
			oldCap = newCap
		}
	}
}

// https://www.cnblogs.com/qcrao-2018/p/10631989.html

// slice的扩容，在1024以下的元素时，都是而被扩容，大于1024时，并不是1.25倍
// 进行内存对齐之后，新 slice 的容量是要 大于等于 老 slice 容量的 2倍或者1.25倍
