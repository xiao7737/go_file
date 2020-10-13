package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7] // 表示长度为4，容量为5，第二个冒号为容量限制

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}

// https://www.cnblogs.com/qcrao-2018/p/10631989.html
/*
[2 3 20]
[4 5 6 7 100 200]
[0 1 2 3 20 5 6 7 100 9]     200加入的时候，s2使用的底层数据不再是slice，所以200没有影响slice
*/

// 多个切片使用的一个底层数组，只有在不发生扩容的情况下，切片元素改变，会影响到原有的slice
