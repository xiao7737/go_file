package main

import "fmt"

func calc(index string, a, b int) int {
	sum := a + b
	fmt.Println(index, a, b, sum)
	return sum
}
func main() {
	a := 1
	b := 2
	defer /*④*/ calc("1", a /*①*/, calc("10", a, b))
	a = 0
	defer /*3*/ calc("2", a /*②*/, calc("20", a, b))
	b = 1
}

//1234执行顺序
/*结果
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4*/
