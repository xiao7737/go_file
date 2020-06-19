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
}
