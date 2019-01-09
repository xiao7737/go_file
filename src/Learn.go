package main

import "fmt"

//函数也是一种类型，申明一个参数类型为函数，就可以将函数作为另一个函数的参数参入
//增强程序的灵活性，用来通用接口
type testInt func(int) bool //申明一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}
func filter(slice []int, f testInt) []int { //参数类型为函数
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice =", slice)
	odd := filter(slice, isOdd)   //函数当作值来传递
	even := filter(slice, isEven) //同上
	fmt.Println("odd value of slice are:", odd)
	fmt.Println("even value of slice are:", even)
}
