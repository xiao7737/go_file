package main

import "fmt"

func main() {
	data := []int{1, 2, 3}
	for _, v := range data {
		v *= 10
	}
	fmt.Println(data)

	for k, v := range data {
		data[k] = 10 * v
	}
	fmt.Println(data)

	data1 := []*struct{ num int }{{1}, {2}, {3}}
	for _, v := range data1 {
		v.num *= 10 // 对于指针类似的值，直接更新，不用采用索引更新
	}
	fmt.Println(data1[0], data1[1], data1[2]) // &{10} &{20} &{30}
}
