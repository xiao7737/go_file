package main

import "fmt"

//从大到小冒泡排序
func main() {
	a := [...]int{1, 3, 4, 5, 7, 9}
	fmt.Println(a)
	length := len(a)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
