package main

import (
	"fmt"
)

func main() {
	fmt.Println(test(true))
}

func test(b bool) (a int) {
	if b {
		a := 2
		a += 1
	}

	return a
}
