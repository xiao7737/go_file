package main

import "fmt"

func main() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[3:6:8]
	fmt.Printf("myslice为 %d, 其容量为: %d，长度为 %d\n", myslice, cap(myslice), len(myslice))
}
